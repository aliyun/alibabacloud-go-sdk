// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDetectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDetectConfigResponseBody
	GetRequestId() *string
}

type DeleteDetectConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BF75EF50-955D-5E1F-AB23-A657C2C6D3C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteDetectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDetectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDetectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDetectConfigResponseBody) SetRequestId(v string) *DeleteDetectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDetectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
