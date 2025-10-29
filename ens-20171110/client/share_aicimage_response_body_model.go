// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareAICImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ShareAICImageResponseBody
	GetRequestId() *string
}

type ShareAICImageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ShareAICImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ShareAICImageResponseBody) GoString() string {
	return s.String()
}

func (s *ShareAICImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ShareAICImageResponseBody) SetRequestId(v string) *ShareAICImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShareAICImageResponseBody) Validate() error {
	return dara.Validate(s)
}
