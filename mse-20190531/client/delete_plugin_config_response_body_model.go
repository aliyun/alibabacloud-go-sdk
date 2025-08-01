// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePluginConfigResponseBody
	GetRequestId() *string
}

type DeletePluginConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC34E4A3-5F1C-4E40-86EA-02EDF967****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePluginConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePluginConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePluginConfigResponseBody) SetRequestId(v string) *DeletePluginConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePluginConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
