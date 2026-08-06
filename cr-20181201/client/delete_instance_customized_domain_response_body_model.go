// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceCustomizedDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstanceCustomizedDomainResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteInstanceCustomizedDomainResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteInstanceCustomizedDomainResponseBody
	GetRequestId() *string
}

type DeleteInstanceCustomizedDomainResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EEE92FA9-3181-5174-8A06-BE2252FA462E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceCustomizedDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceCustomizedDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceCustomizedDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceCustomizedDomainResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteInstanceCustomizedDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceCustomizedDomainResponseBody) SetCode(v string) *DeleteInstanceCustomizedDomainResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceCustomizedDomainResponseBody) SetIsSuccess(v bool) *DeleteInstanceCustomizedDomainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteInstanceCustomizedDomainResponseBody) SetRequestId(v string) *DeleteInstanceCustomizedDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceCustomizedDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
