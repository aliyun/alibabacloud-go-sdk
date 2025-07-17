// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArmsAgentDownloadUrl(v string) *GetAgentDownloadUrlResponseBody
	GetArmsAgentDownloadUrl() *string
	SetRequestId(v string) *GetAgentDownloadUrlResponseBody
	GetRequestId() *string
}

type GetAgentDownloadUrlResponseBody struct {
	// The download URL of the ARMS agent.
	//
	// example:
	//
	// http://arms-apm-hangzhou.oss-cn-hangzhou-internal.aliyuncs.com/2.7.1.1/
	ArmsAgentDownloadUrl *string `json:"ArmsAgentDownloadUrl,omitempty" xml:"ArmsAgentDownloadUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14043452-D486-4EA1-80C9-BA73FB81****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlResponseBody) GetArmsAgentDownloadUrl() *string {
	return s.ArmsAgentDownloadUrl
}

func (s *GetAgentDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentDownloadUrlResponseBody) SetArmsAgentDownloadUrl(v string) *GetAgentDownloadUrlResponseBody {
	s.ArmsAgentDownloadUrl = &v
	return s
}

func (s *GetAgentDownloadUrlResponseBody) SetRequestId(v string) *GetAgentDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
