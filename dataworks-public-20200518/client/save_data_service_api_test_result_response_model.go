// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDataServiceApiTestResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveDataServiceApiTestResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveDataServiceApiTestResultResponse
	GetStatusCode() *int32
	SetBody(v *SaveDataServiceApiTestResultResponseBody) *SaveDataServiceApiTestResultResponse
	GetBody() *SaveDataServiceApiTestResultResponseBody
}

type SaveDataServiceApiTestResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveDataServiceApiTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveDataServiceApiTestResultResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveDataServiceApiTestResultResponse) GoString() string {
	return s.String()
}

func (s *SaveDataServiceApiTestResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveDataServiceApiTestResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveDataServiceApiTestResultResponse) GetBody() *SaveDataServiceApiTestResultResponseBody {
	return s.Body
}

func (s *SaveDataServiceApiTestResultResponse) SetHeaders(v map[string]*string) *SaveDataServiceApiTestResultResponse {
	s.Headers = v
	return s
}

func (s *SaveDataServiceApiTestResultResponse) SetStatusCode(v int32) *SaveDataServiceApiTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveDataServiceApiTestResultResponse) SetBody(v *SaveDataServiceApiTestResultResponseBody) *SaveDataServiceApiTestResultResponse {
	s.Body = v
	return s
}

func (s *SaveDataServiceApiTestResultResponse) Validate() error {
	return dara.Validate(s)
}
