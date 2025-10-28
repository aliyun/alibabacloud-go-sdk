// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RestartApplicationRequest
	GetAppId() *string
	SetEccInfo(v string) *RestartApplicationRequest
	GetEccInfo() *string
}

type RestartApplicationRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// c627c157-560d*******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the elastic compute container (ECC) that corresponds to the ECS instance on which you want to restart the application. You can call the QueryApplicationStatus operation to query the ECC ID. For more information, see [QueryApplicationStatus](https://help.aliyun.com/document_detail/149394.html).
	//
	// 	- Separate multiple ECC IDs with commas (,).
	//
	// 	- If you leave this parameter empty, the application will be restarted on all the ECS instances deployed with the application.
	//
	// example:
	//
	// 006c0ea5-5f8d-4398-af1e-**********
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
}

func (s RestartApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartApplicationRequest) GoString() string {
	return s.String()
}

func (s *RestartApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *RestartApplicationRequest) GetEccInfo() *string {
	return s.EccInfo
}

func (s *RestartApplicationRequest) SetAppId(v string) *RestartApplicationRequest {
	s.AppId = &v
	return s
}

func (s *RestartApplicationRequest) SetEccInfo(v string) *RestartApplicationRequest {
	s.EccInfo = &v
	return s
}

func (s *RestartApplicationRequest) Validate() error {
	return dara.Validate(s)
}
