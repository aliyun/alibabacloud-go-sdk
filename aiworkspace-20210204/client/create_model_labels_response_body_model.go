// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateModelLabelsResponseBody
	GetRequestId() *string
}

type CreateModelLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F81D9EC0-1872-50F5-A96C-A0647D****1D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelLabelsResponseBody) SetRequestId(v string) *CreateModelLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
