// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *GetAppInfosRequest
	GetAppIds() *string
}

type GetAppInfosRequest struct {
	// The IDs of applications. You can obtain application IDs from the response to the [CreateAppInfo](https://help.aliyun.com/document_detail/113266.html) or [ListAppInfo](https://help.aliyun.com/document_detail/114000.html) operation.
	//
	// 	- You can specify a maximum of 10 application IDs.
	//
	// 	- Separate application IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// app-****
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
}

func (s GetAppInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppInfosRequest) GoString() string {
	return s.String()
}

func (s *GetAppInfosRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *GetAppInfosRequest) SetAppIds(v string) *GetAppInfosRequest {
	s.AppIds = &v
	return s
}

func (s *GetAppInfosRequest) Validate() error {
	return dara.Validate(s)
}
