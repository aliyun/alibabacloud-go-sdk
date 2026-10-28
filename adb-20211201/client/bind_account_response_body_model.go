// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindAccountResponseBody
	GetRequestId() *string
}

type BindAccountResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DFF27323-3868-5F8A-917D-5D1D06B6BC0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *BindAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindAccountResponseBody) SetRequestId(v string) *BindAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
