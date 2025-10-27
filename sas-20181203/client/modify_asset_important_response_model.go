// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAssetImportantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAssetImportantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAssetImportantResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAssetImportantResponseBody) *ModifyAssetImportantResponse
	GetBody() *ModifyAssetImportantResponseBody
}

type ModifyAssetImportantResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAssetImportantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAssetImportantResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetImportantResponse) GoString() string {
	return s.String()
}

func (s *ModifyAssetImportantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAssetImportantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAssetImportantResponse) GetBody() *ModifyAssetImportantResponseBody {
	return s.Body
}

func (s *ModifyAssetImportantResponse) SetHeaders(v map[string]*string) *ModifyAssetImportantResponse {
	s.Headers = v
	return s
}

func (s *ModifyAssetImportantResponse) SetStatusCode(v int32) *ModifyAssetImportantResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAssetImportantResponse) SetBody(v *ModifyAssetImportantResponseBody) *ModifyAssetImportantResponse {
	s.Body = v
	return s
}

func (s *ModifyAssetImportantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
