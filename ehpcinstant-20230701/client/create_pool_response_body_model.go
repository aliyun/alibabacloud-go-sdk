// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePoolResponseBody
	GetRequestId() *string
}

type CreatePoolResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePoolResponseBody) SetRequestId(v string) *CreatePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
