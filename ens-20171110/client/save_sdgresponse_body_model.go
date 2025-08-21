// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSDGResponseBody
	GetRequestId() *string
}

type SaveSDGResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSDGResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSDGResponseBody) SetRequestId(v string) *SaveSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSDGResponseBody) Validate() error {
	return dara.Validate(s)
}
