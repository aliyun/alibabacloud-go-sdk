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
	// Id of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// 0B6F0F99-EB17-51D5-AAC8-AD78A26E18DD
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
