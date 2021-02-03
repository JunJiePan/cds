package handler

import (
	"net/http"

	logic2 "github.com/tal-tech/cds/galaxy/internal/logic"
	"github.com/tal-tech/cds/galaxy/internal/svc"
	"github.com/tal-tech/cds/galaxy/internal/types"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func generateCreateSqlHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic2.NewGenerateCreateSqlLogic(r.Context(), ctx)
		var req types.GenerateCreateSqlRequest
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := l.GenerateCreateSql(req)
		formatFullResponse(resp, err, w, r)
	}
}
