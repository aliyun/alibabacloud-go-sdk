// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateParameterResponseBody
	GetRequestId() *string
	SetVersion(v int32) *UpdateParameterResponseBody
	GetVersion() *int32
}

type UpdateParameterResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateParameterResponseBody) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateParameterResponseBody) SetRequestId(v string) *UpdateParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateParameterResponseBody) SetVersion(v int32) *UpdateParameterResponseBody {
	s.Version = &v
	return s
}

func (s *UpdateParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
