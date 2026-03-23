// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddApListToApgroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApGroupId(v string) *AddApListToApgroupRequest
	GetApGroupId() *string
	SetApMacList(v map[string]interface{}) *AddApListToApgroupRequest
	GetApMacList() map[string]interface{}
	SetAppCode(v string) *AddApListToApgroupRequest
	GetAppCode() *string
	SetAppName(v string) *AddApListToApgroupRequest
	GetAppName() *string
}

type AddApListToApgroupRequest struct {
	// This parameter is required.
	ApGroupId *string `json:"ApGroupId,omitempty" xml:"ApGroupId,omitempty"`
	// This parameter is required.
	ApMacList map[string]interface{} `json:"ApMacList,omitempty" xml:"ApMacList,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s AddApListToApgroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddApListToApgroupRequest) GoString() string {
	return s.String()
}

func (s *AddApListToApgroupRequest) GetApGroupId() *string {
	return s.ApGroupId
}

func (s *AddApListToApgroupRequest) GetApMacList() map[string]interface{} {
	return s.ApMacList
}

func (s *AddApListToApgroupRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *AddApListToApgroupRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddApListToApgroupRequest) SetApGroupId(v string) *AddApListToApgroupRequest {
	s.ApGroupId = &v
	return s
}

func (s *AddApListToApgroupRequest) SetApMacList(v map[string]interface{}) *AddApListToApgroupRequest {
	s.ApMacList = v
	return s
}

func (s *AddApListToApgroupRequest) SetAppCode(v string) *AddApListToApgroupRequest {
	s.AppCode = &v
	return s
}

func (s *AddApListToApgroupRequest) SetAppName(v string) *AddApListToApgroupRequest {
	s.AppName = &v
	return s
}

func (s *AddApListToApgroupRequest) Validate() error {
	return dara.Validate(s)
}
