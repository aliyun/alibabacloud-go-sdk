// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadModuleResponseBody
	GetRequestId() *string
	SetVersion(v string) *UploadModuleResponseBody
	GetVersion() *string
}

type UploadModuleResponseBody struct {
	// example:
	//
	// 0DDD8773-5756-5508-BE36-D03DE43E2450
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Version   *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UploadModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadModuleResponseBody) GoString() string {
	return s.String()
}

func (s *UploadModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadModuleResponseBody) GetVersion() *string {
	return s.Version
}

func (s *UploadModuleResponseBody) SetRequestId(v string) *UploadModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadModuleResponseBody) SetVersion(v string) *UploadModuleResponseBody {
	s.Version = &v
	return s
}

func (s *UploadModuleResponseBody) Validate() error {
	return dara.Validate(s)
}
