// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCollectionDataMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCollectionDataMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCollectionDataMetadataResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCollectionDataMetadataResponseBody) *UpdateCollectionDataMetadataResponse
	GetBody() *UpdateCollectionDataMetadataResponseBody
}

type UpdateCollectionDataMetadataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCollectionDataMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCollectionDataMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectionDataMetadataResponse) GoString() string {
	return s.String()
}

func (s *UpdateCollectionDataMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCollectionDataMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCollectionDataMetadataResponse) GetBody() *UpdateCollectionDataMetadataResponseBody {
	return s.Body
}

func (s *UpdateCollectionDataMetadataResponse) SetHeaders(v map[string]*string) *UpdateCollectionDataMetadataResponse {
	s.Headers = v
	return s
}

func (s *UpdateCollectionDataMetadataResponse) SetStatusCode(v int32) *UpdateCollectionDataMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCollectionDataMetadataResponse) SetBody(v *UpdateCollectionDataMetadataResponseBody) *UpdateCollectionDataMetadataResponse {
	s.Body = v
	return s
}

func (s *UpdateCollectionDataMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
