// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCloudAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UploadCloudAppResponseBody
	GetAppId() *string
	SetRequestId(v string) *UploadCloudAppResponseBody
	GetRequestId() *string
}

type UploadCloudAppResponseBody struct {
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadCloudAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadCloudAppResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCloudAppResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *UploadCloudAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadCloudAppResponseBody) SetAppId(v string) *UploadCloudAppResponseBody {
	s.AppId = &v
	return s
}

func (s *UploadCloudAppResponseBody) SetRequestId(v string) *UploadCloudAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadCloudAppResponseBody) Validate() error {
	return dara.Validate(s)
}
