// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApiProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachApiProductResponseBody
	GetRequestId() *string
}

type AttachApiProductResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BA20890E-75C7-41BC-9C8B-73276B58F550
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachApiProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachApiProductResponseBody) GoString() string {
	return s.String()
}

func (s *AttachApiProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachApiProductResponseBody) SetRequestId(v string) *AttachApiProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachApiProductResponseBody) Validate() error {
	return dara.Validate(s)
}
