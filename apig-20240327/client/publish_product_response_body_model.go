// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishProductResponseBody
	GetCode() *string
	SetData(v *PublishProductResponseBodyData) *PublishProductResponseBody
	GetData() *PublishProductResponseBodyData
	SetMessage(v string) *PublishProductResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishProductResponseBody
	GetRequestId() *string
}

type PublishProductResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *PublishProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PublishProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishProductResponseBody) GoString() string {
	return s.String()
}

func (s *PublishProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishProductResponseBody) GetData() *PublishProductResponseBodyData {
	return s.Data
}

func (s *PublishProductResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishProductResponseBody) SetCode(v string) *PublishProductResponseBody {
	s.Code = &v
	return s
}

func (s *PublishProductResponseBody) SetData(v *PublishProductResponseBodyData) *PublishProductResponseBody {
	s.Data = v
	return s
}

func (s *PublishProductResponseBody) SetMessage(v string) *PublishProductResponseBody {
	s.Message = &v
	return s
}

func (s *PublishProductResponseBody) SetRequestId(v string) *PublishProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishProductResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishProductResponseBodyData struct {
	// The publication record ID, which identifies the publication relationship between the published API product and the target portal. You can query this ID in the response of ListProductPublications at /himarket/v1/products/{productId}/publications. You can also use this ID as the publicationId path parameter in UnpublishProduct to unpublish the record.
	//
	// example:
	//
	// publication-xxxxxxxx
	PublicationId *string `json:"publicationId,omitempty" xml:"publicationId,omitempty"`
}

func (s PublishProductResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PublishProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *PublishProductResponseBodyData) GetPublicationId() *string {
	return s.PublicationId
}

func (s *PublishProductResponseBodyData) SetPublicationId(v string) *PublishProductResponseBodyData {
	s.PublicationId = &v
	return s
}

func (s *PublishProductResponseBodyData) Validate() error {
	return dara.Validate(s)
}
