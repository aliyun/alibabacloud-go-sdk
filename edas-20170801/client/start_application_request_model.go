// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartApplicationRequest
	GetAppId() *string
	SetEccInfo(v string) *StartApplicationRequest
	GetEccInfo() *string
}

type StartApplicationRequest struct {
	// The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413**********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the elastic compute container (ECC) that corresponds to the Elastic Compute Service (ECS) instance on which you want to start the application. Separate multiple ECC IDs with commas (,). You can call the QueryApplicationStatus operation to query the ECC ID. For more information, see [QueryApplicationStatus](https://help.aliyun.com/document_detail/149394.html).
	//
	// example:
	//
	// ""
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
}

func (s StartApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartApplicationRequest) GoString() string {
	return s.String()
}

func (s *StartApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartApplicationRequest) GetEccInfo() *string {
	return s.EccInfo
}

func (s *StartApplicationRequest) SetAppId(v string) *StartApplicationRequest {
	s.AppId = &v
	return s
}

func (s *StartApplicationRequest) SetEccInfo(v string) *StartApplicationRequest {
	s.EccInfo = &v
	return s
}

func (s *StartApplicationRequest) Validate() error {
	return dara.Validate(s)
}
