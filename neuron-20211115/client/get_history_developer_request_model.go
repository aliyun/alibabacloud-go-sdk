// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoryDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *GetHistoryDeveloperRequest
	GetEnterpriseId() *int64
}

type GetHistoryDeveloperRequest struct {
	EnterpriseId *int64 `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
}

func (s GetHistoryDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryDeveloperRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryDeveloperRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *GetHistoryDeveloperRequest) SetEnterpriseId(v int64) *GetHistoryDeveloperRequest {
	s.EnterpriseId = &v
	return s
}

func (s *GetHistoryDeveloperRequest) Validate() error {
	return dara.Validate(s)
}
