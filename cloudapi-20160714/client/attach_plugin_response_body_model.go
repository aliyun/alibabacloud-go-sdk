// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachPluginResponseBody
	GetRequestId() *string
}

type AttachPluginResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachPluginResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachPluginResponseBody) SetRequestId(v string) *AttachPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
