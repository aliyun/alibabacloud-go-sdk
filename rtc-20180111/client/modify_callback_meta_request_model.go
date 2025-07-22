// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCallbackMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyCallbackMetaRequest
	GetAppId() *string
	SetCallback(v *ModifyCallbackMetaRequestCallback) *ModifyCallbackMetaRequest
	GetCallback() *ModifyCallbackMetaRequestCallback
	SetOwnerId(v int64) *ModifyCallbackMetaRequest
	GetOwnerId() *int64
}

type ModifyCallbackMetaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223***JQb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Callback *ModifyCallbackMetaRequestCallback `json:"Callback,omitempty" xml:"Callback,omitempty" type:"Struct"`
	OwnerId  *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyCallbackMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCallbackMetaRequest) GoString() string {
	return s.String()
}

func (s *ModifyCallbackMetaRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyCallbackMetaRequest) GetCallback() *ModifyCallbackMetaRequestCallback {
	return s.Callback
}

func (s *ModifyCallbackMetaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCallbackMetaRequest) SetAppId(v string) *ModifyCallbackMetaRequest {
	s.AppId = &v
	return s
}

func (s *ModifyCallbackMetaRequest) SetCallback(v *ModifyCallbackMetaRequestCallback) *ModifyCallbackMetaRequest {
	s.Callback = v
	return s
}

func (s *ModifyCallbackMetaRequest) SetOwnerId(v int64) *ModifyCallbackMetaRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCallbackMetaRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyCallbackMetaRequestCallback struct {
	// This parameter is required.
	//
	// example:
	//
	// RecordEvent
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyun.com
	Conf     *string  `json:"Conf,omitempty" xml:"Conf,omitempty"`
	SubEvent []*int32 `json:"SubEvent,omitempty" xml:"SubEvent,omitempty" type:"Repeated"`
}

func (s ModifyCallbackMetaRequestCallback) String() string {
	return dara.Prettify(s)
}

func (s ModifyCallbackMetaRequestCallback) GoString() string {
	return s.String()
}

func (s *ModifyCallbackMetaRequestCallback) GetCategory() *string {
	return s.Category
}

func (s *ModifyCallbackMetaRequestCallback) GetConf() *string {
	return s.Conf
}

func (s *ModifyCallbackMetaRequestCallback) GetSubEvent() []*int32 {
	return s.SubEvent
}

func (s *ModifyCallbackMetaRequestCallback) SetCategory(v string) *ModifyCallbackMetaRequestCallback {
	s.Category = &v
	return s
}

func (s *ModifyCallbackMetaRequestCallback) SetConf(v string) *ModifyCallbackMetaRequestCallback {
	s.Conf = &v
	return s
}

func (s *ModifyCallbackMetaRequestCallback) SetSubEvent(v []*int32) *ModifyCallbackMetaRequestCallback {
	s.SubEvent = v
	return s
}

func (s *ModifyCallbackMetaRequestCallback) Validate() error {
	return dara.Validate(s)
}
