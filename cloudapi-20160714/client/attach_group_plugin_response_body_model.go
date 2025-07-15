// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachGroupPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachGroupPluginResponseBody
	GetRequestId() *string
}

type AttachGroupPluginResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D6E46F10-F26C-4AA0-BB69-FE2743D9AE62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachGroupPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachGroupPluginResponseBody) GoString() string {
	return s.String()
}

func (s *AttachGroupPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachGroupPluginResponseBody) SetRequestId(v string) *AttachGroupPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachGroupPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
