// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGdnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteGdnInstanceResponseBodyData) *DeleteGdnInstanceResponseBody
	GetData() *DeleteGdnInstanceResponseBodyData
	SetMessage(v string) *DeleteGdnInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGdnInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGdnInstanceResponseBody
	GetSuccess() *bool
}

type DeleteGdnInstanceResponseBody struct {
	Data *DeleteGdnInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 14036EBE-CF2E-44DB-ACE9-AC6157D3A6FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGdnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGdnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGdnInstanceResponseBody) GetData() *DeleteGdnInstanceResponseBodyData {
	return s.Data
}

func (s *DeleteGdnInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGdnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGdnInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGdnInstanceResponseBody) SetData(v *DeleteGdnInstanceResponseBodyData) *DeleteGdnInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteGdnInstanceResponseBody) SetMessage(v string) *DeleteGdnInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGdnInstanceResponseBody) SetRequestId(v string) *DeleteGdnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGdnInstanceResponseBody) SetSuccess(v bool) *DeleteGdnInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGdnInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteGdnInstanceResponseBodyData struct {
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteGdnInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteGdnInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteGdnInstanceResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DeleteGdnInstanceResponseBodyData) SetTaskId(v int32) *DeleteGdnInstanceResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DeleteGdnInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
