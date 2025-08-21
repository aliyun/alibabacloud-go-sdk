// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ImportImageResponseBody
	GetImageId() *string
	SetRequestId(v string) *ImportImageResponseBody
	GetRequestId() *string
}

type ImportImageResponseBody struct {
	// example:
	//
	// m-5wn1dh*******b48f440ntvad
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportImageResponseBody) GoString() string {
	return s.String()
}

func (s *ImportImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *ImportImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportImageResponseBody) SetImageId(v string) *ImportImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *ImportImageResponseBody) SetRequestId(v string) *ImportImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportImageResponseBody) Validate() error {
	return dara.Validate(s)
}
