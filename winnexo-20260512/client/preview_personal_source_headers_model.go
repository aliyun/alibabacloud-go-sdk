// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewPersonalSourceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PreviewPersonalSourceHeaders
	GetCommonHeaders() map[string]*string
	SetRequestId(v string) *PreviewPersonalSourceHeaders
	GetRequestId() *string
}

type PreviewPersonalSourceHeaders struct {
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

func (s PreviewPersonalSourceHeaders) String() string {
	return dara.Prettify(s)
}

func (s PreviewPersonalSourceHeaders) GoString() string {
	return s.String()
}

func (s *PreviewPersonalSourceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PreviewPersonalSourceHeaders) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewPersonalSourceHeaders) SetCommonHeaders(v map[string]*string) *PreviewPersonalSourceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PreviewPersonalSourceHeaders) SetRequestId(v string) *PreviewPersonalSourceHeaders {
	s.RequestId = &v
	return s
}

func (s *PreviewPersonalSourceHeaders) Validate() error {
	return dara.Validate(s)
}
