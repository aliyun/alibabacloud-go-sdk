// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAyncTradeDocumentPackageExtractSmartAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AyncTradeDocumentPackageExtractSmartAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AyncTradeDocumentPackageExtractSmartAppResponse
	GetStatusCode() *int32
	SetBody(v *AyncTradeDocumentPackageExtractSmartAppResponseBody) *AyncTradeDocumentPackageExtractSmartAppResponse
	GetBody() *AyncTradeDocumentPackageExtractSmartAppResponseBody
}

type AyncTradeDocumentPackageExtractSmartAppResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AyncTradeDocumentPackageExtractSmartAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AyncTradeDocumentPackageExtractSmartAppResponse) String() string {
	return dara.Prettify(s)
}

func (s AyncTradeDocumentPackageExtractSmartAppResponse) GoString() string {
	return s.String()
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) GetBody() *AyncTradeDocumentPackageExtractSmartAppResponseBody {
	return s.Body
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) SetHeaders(v map[string]*string) *AyncTradeDocumentPackageExtractSmartAppResponse {
	s.Headers = v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) SetStatusCode(v int32) *AyncTradeDocumentPackageExtractSmartAppResponse {
	s.StatusCode = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) SetBody(v *AyncTradeDocumentPackageExtractSmartAppResponseBody) *AyncTradeDocumentPackageExtractSmartAppResponse {
	s.Body = v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppResponse) Validate() error {
	return dara.Validate(s)
}
