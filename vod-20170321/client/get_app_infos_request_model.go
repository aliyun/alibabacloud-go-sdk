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
	// The list of application IDs. The list consists of one or more application IDs. An application ID is the value of the AppId parameter returned by the [CreateAppInfo](https://help.aliyun.com/document_detail/113266.html) or [GetAppInfos](https://help.aliyun.com/document_detail/114000.html) operation.
	//
	// - A maximum of 10 IDs are supported.
	//
	// - Separate multiple IDs with commas (,).
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
