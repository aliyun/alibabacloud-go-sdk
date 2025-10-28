// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteApplicationRequest
	GetAppId() *string
}

type DeleteApplicationRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-*********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteApplicationRequest) SetAppId(v string) *DeleteApplicationRequest {
	s.AppId = &v
	return s
}

func (s *DeleteApplicationRequest) Validate() error {
	return dara.Validate(s)
}
