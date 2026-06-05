// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSupabaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppSupabaseInstanceRequest
	GetBizId() *string
}

type GetAppSupabaseInstanceRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
}

func (s GetAppSupabaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseInstanceRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppSupabaseInstanceRequest) SetBizId(v string) *GetAppSupabaseInstanceRequest {
	s.BizId = &v
	return s
}

func (s *GetAppSupabaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
