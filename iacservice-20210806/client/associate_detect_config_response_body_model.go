// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateDetectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateDetectConfigResponseBody
	GetRequestId() *string
}

type AssociateDetectConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BF75EF50-955D-5E1F-AB23-A657C2C6D3C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AssociateDetectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateDetectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateDetectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateDetectConfigResponseBody) SetRequestId(v string) *AssociateDetectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateDetectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
