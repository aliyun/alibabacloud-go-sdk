// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudGroupExpansionServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridCloudGroupExpansionServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridCloudGroupExpansionServerResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridCloudGroupExpansionServerResponseBody) *ModifyHybridCloudGroupExpansionServerResponse
	GetBody() *ModifyHybridCloudGroupExpansionServerResponseBody
}

type ModifyHybridCloudGroupExpansionServerResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridCloudGroupExpansionServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridCloudGroupExpansionServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudGroupExpansionServerResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudGroupExpansionServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridCloudGroupExpansionServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridCloudGroupExpansionServerResponse) GetBody() *ModifyHybridCloudGroupExpansionServerResponseBody {
	return s.Body
}

func (s *ModifyHybridCloudGroupExpansionServerResponse) SetHeaders(v map[string]*string) *ModifyHybridCloudGroupExpansionServerResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerResponse) SetStatusCode(v int32) *ModifyHybridCloudGroupExpansionServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerResponse) SetBody(v *ModifyHybridCloudGroupExpansionServerResponseBody) *ModifyHybridCloudGroupExpansionServerResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerResponse) Validate() error {
	return dara.Validate(s)
}
