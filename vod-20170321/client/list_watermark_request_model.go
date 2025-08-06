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
	// The ID of the application. Default value: **app-1000000**.
	//
	// If you have activated the multi-application service, specify the ID of the application to query all image and text watermark templates in the specified application. If you leave this parameter empty, image and text watermark templates in all applications are queried. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
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
