// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListWatermarkRequest
	GetAppId() *string
}

type ListWatermarkRequest struct {
	// The application ID. Default value: **app-1000000**.
	//
	// If the multi-application service is enabled, you can specify an application ID to query watermark templates under the specified application. If you do not specify this parameter, watermark templates under all applications are returned. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWatermarkRequest) GoString() string {
	return s.String()
}

func (s *ListWatermarkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListWatermarkRequest) SetAppId(v string) *ListWatermarkRequest {
	s.AppId = &v
	return s
}

func (s *ListWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
