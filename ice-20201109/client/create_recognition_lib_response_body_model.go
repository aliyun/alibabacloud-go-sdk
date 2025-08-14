// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecognitionLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *CreateRecognitionLibResponseBody
	GetLibId() *string
	SetRequestId(v string) *CreateRecognitionLibResponseBody
	GetRequestId() *string
}

type CreateRecognitionLibResponseBody struct {
	// example:
	//
	// *************24b47865c6**************
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecognitionLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecognitionLibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecognitionLibResponseBody) GetLibId() *string {
	return s.LibId
}

func (s *CreateRecognitionLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecognitionLibResponseBody) SetLibId(v string) *CreateRecognitionLibResponseBody {
	s.LibId = &v
	return s
}

func (s *CreateRecognitionLibResponseBody) SetRequestId(v string) *CreateRecognitionLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecognitionLibResponseBody) Validate() error {
	return dara.Validate(s)
}
