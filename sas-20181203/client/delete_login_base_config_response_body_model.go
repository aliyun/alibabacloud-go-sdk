// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoginBaseConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLoginBaseConfigResponseBody
	GetRequestId() *string
}

type DeleteLoginBaseConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoginBaseConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoginBaseConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoginBaseConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLoginBaseConfigResponseBody) SetRequestId(v string) *DeleteLoginBaseConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLoginBaseConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
