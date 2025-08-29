// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParamId(v int64) *CreateParamResponseBody
	GetParamId() *int64
	SetRequestId(v string) *CreateParamResponseBody
	GetRequestId() *string
}

type CreateParamResponseBody struct {
	// example:
	//
	// 4
	ParamId *int64 `json:"ParamId,omitempty" xml:"ParamId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateParamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParamResponseBody) GetParamId() *int64 {
	return s.ParamId
}

func (s *CreateParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateParamResponseBody) SetParamId(v int64) *CreateParamResponseBody {
	s.ParamId = &v
	return s
}

func (s *CreateParamResponseBody) SetRequestId(v string) *CreateParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateParamResponseBody) Validate() error {
	return dara.Validate(s)
}
