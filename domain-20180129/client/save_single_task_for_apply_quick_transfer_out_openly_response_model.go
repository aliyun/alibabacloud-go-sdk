// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForApplyQuickTransferOutOpenlyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse
	GetBody() *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody
}

type SaveSingleTaskForApplyQuickTransferOutOpenlyResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForApplyQuickTransferOutOpenlyResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForApplyQuickTransferOutOpenlyResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse) GetBody() *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse) SetStatusCode(v int32) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse) SetBody(v *SaveSingleTaskForApplyQuickTransferOutOpenlyResponseBody) *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForApplyQuickTransferOutOpenlyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
