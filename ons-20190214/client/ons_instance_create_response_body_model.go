// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsInstanceCreateResponseBodyData) *OnsInstanceCreateResponseBody
	GetData() *OnsInstanceCreateResponseBodyData
	SetRequestId(v string) *OnsInstanceCreateResponseBody
	GetRequestId() *string
}

type OnsInstanceCreateResponseBody struct {
	// The result returned.
	Data *OnsInstanceCreateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07E3902-B92E-44A6-B6C5-6AA111111****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsInstanceCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceCreateResponseBody) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateResponseBody) GetData() *OnsInstanceCreateResponseBodyData {
	return s.Data
}

func (s *OnsInstanceCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsInstanceCreateResponseBody) SetData(v *OnsInstanceCreateResponseBodyData) *OnsInstanceCreateResponseBody {
	s.Data = v
	return s
}

func (s *OnsInstanceCreateResponseBody) SetRequestId(v string) *OnsInstanceCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsInstanceCreateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsInstanceCreateResponseBodyData struct {
	// The ID of the instance that you created.
	//
	// example:
	//
	// MQ_INST_188077086902****_BXSuW61e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The edition of the instance that you created. Valid value:
	//
	// 	- **1**: Standard Edition instances
	//
	// example:
	//
	// 1
	InstanceType *int32 `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s OnsInstanceCreateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceCreateResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsInstanceCreateResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsInstanceCreateResponseBodyData) GetInstanceType() *int32 {
	return s.InstanceType
}

func (s *OnsInstanceCreateResponseBodyData) SetInstanceId(v string) *OnsInstanceCreateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsInstanceCreateResponseBodyData) SetInstanceType(v int32) *OnsInstanceCreateResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *OnsInstanceCreateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
