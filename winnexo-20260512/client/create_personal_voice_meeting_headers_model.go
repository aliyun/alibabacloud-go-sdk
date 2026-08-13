// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalVoiceMeetingHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreatePersonalVoiceMeetingHeaders
	GetCommonHeaders() map[string]*string
	SetRequestId(v string) *CreatePersonalVoiceMeetingHeaders
	GetRequestId() *string
}

type CreatePersonalVoiceMeetingHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// 请求追踪 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePersonalVoiceMeetingHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalVoiceMeetingHeaders) GoString() string {
	return s.String()
}

func (s *CreatePersonalVoiceMeetingHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreatePersonalVoiceMeetingHeaders) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalVoiceMeetingHeaders) SetCommonHeaders(v map[string]*string) *CreatePersonalVoiceMeetingHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePersonalVoiceMeetingHeaders) SetRequestId(v string) *CreatePersonalVoiceMeetingHeaders {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingHeaders) Validate() error {
	return dara.Validate(s)
}
