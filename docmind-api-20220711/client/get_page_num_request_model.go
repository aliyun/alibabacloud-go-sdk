// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageNumRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetPageNumRequest
	GetBizId() *string
}

type GetPageNumRequest struct {
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GetPageNumRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPageNumRequest) GoString() string {
	return s.String()
}

func (s *GetPageNumRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetPageNumRequest) SetBizId(v string) *GetPageNumRequest {
	s.BizId = &v
	return s
}

func (s *GetPageNumRequest) Validate() error {
	return dara.Validate(s)
}
