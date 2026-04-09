// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteParameterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteParameterResponseBody
	GetSuccess() *bool
}

type DeleteParameterResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteParameterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteParameterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteParameterResponseBody) SetRequestId(v string) *DeleteParameterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteParameterResponseBody) SetSuccess(v bool) *DeleteParameterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteParameterResponseBody) Validate() error {
	return dara.Validate(s)
}
