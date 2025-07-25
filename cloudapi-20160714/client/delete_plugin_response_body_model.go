// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePluginResponseBody
	GetRequestId() *string
}

type DeletePluginResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePluginResponseBody) SetRequestId(v string) *DeletePluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePluginResponseBody) Validate() error {
	return dara.Validate(s)
}
