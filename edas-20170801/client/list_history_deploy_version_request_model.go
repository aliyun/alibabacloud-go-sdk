// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoryDeployVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListHistoryDeployVersionRequest
	GetAppId() *string
}

type ListHistoryDeployVersionRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListHistoryDeployVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHistoryDeployVersionRequest) GoString() string {
	return s.String()
}

func (s *ListHistoryDeployVersionRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListHistoryDeployVersionRequest) SetAppId(v string) *ListHistoryDeployVersionRequest {
	s.AppId = &v
	return s
}

func (s *ListHistoryDeployVersionRequest) Validate() error {
	return dara.Validate(s)
}
