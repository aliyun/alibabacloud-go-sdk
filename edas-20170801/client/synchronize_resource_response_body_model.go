// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SynchronizeResourceResponseBody
	GetCode() *int32
	SetData(v string) *SynchronizeResourceResponseBody
	GetData() *string
	SetMessage(v string) *SynchronizeResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SynchronizeResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SynchronizeResourceResponseBody
	GetSuccess() *bool
}

type SynchronizeResourceResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned for the request.
	//
	// example:
	//
	// PopSyncResource success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F8DFGED-K98***************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the resources are synchronized. Valid values:
	//
	// 	- **true**: The resources are synchronized.
	//
	// 	- **false**: The resources fail to be synchronized.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SynchronizeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SynchronizeResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SynchronizeResourceResponseBody) GetData() *string {
	return s.Data
}

func (s *SynchronizeResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SynchronizeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SynchronizeResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SynchronizeResourceResponseBody) SetCode(v int32) *SynchronizeResourceResponseBody {
	s.Code = &v
	return s
}

func (s *SynchronizeResourceResponseBody) SetData(v string) *SynchronizeResourceResponseBody {
	s.Data = &v
	return s
}

func (s *SynchronizeResourceResponseBody) SetMessage(v string) *SynchronizeResourceResponseBody {
	s.Message = &v
	return s
}

func (s *SynchronizeResourceResponseBody) SetRequestId(v string) *SynchronizeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SynchronizeResourceResponseBody) SetSuccess(v bool) *SynchronizeResourceResponseBody {
	s.Success = &v
	return s
}

func (s *SynchronizeResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
