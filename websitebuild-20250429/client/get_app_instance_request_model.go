// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppInstanceRequest
	GetBizId() *string
}

type GetAppInstanceRequest struct {
	// example:
	//
	// WS20250801152639000005
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GetAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetAppInstanceRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceRequest) SetBizId(v string) *GetAppInstanceRequest {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceRequest) Validate() error {
	return dara.Validate(s)
}
