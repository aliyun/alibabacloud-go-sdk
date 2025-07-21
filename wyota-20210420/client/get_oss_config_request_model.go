// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v int32) *GetOssConfigRequest
	GetType() *int32
}

type GetOssConfigRequest struct {
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetOssConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssConfigRequest) GoString() string {
	return s.String()
}

func (s *GetOssConfigRequest) GetType() *int32 {
	return s.Type
}

func (s *GetOssConfigRequest) SetType(v int32) *GetOssConfigRequest {
	s.Type = &v
	return s
}

func (s *GetOssConfigRequest) Validate() error {
	return dara.Validate(s)
}
