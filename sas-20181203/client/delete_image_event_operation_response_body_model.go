// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageEventOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteImageEventOperationResponseBody
	GetCode() *string
	SetData(v *DeleteImageEventOperationResponseBodyData) *DeleteImageEventOperationResponseBody
	GetData() *DeleteImageEventOperationResponseBodyData
	SetMessage(v string) *DeleteImageEventOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteImageEventOperationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteImageEventOperationResponseBody
	GetSuccess() *bool
}

type DeleteImageEventOperationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteImageEventOperationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3D7C47D-3F11-57BB-90E8-E5C20C61****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageEventOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageEventOperationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageEventOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteImageEventOperationResponseBody) GetData() *DeleteImageEventOperationResponseBodyData {
	return s.Data
}

func (s *DeleteImageEventOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteImageEventOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageEventOperationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteImageEventOperationResponseBody) SetCode(v string) *DeleteImageEventOperationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageEventOperationResponseBody) SetData(v *DeleteImageEventOperationResponseBodyData) *DeleteImageEventOperationResponseBody {
	s.Data = v
	return s
}

func (s *DeleteImageEventOperationResponseBody) SetMessage(v string) *DeleteImageEventOperationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteImageEventOperationResponseBody) SetRequestId(v string) *DeleteImageEventOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageEventOperationResponseBody) SetSuccess(v bool) *DeleteImageEventOperationResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteImageEventOperationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteImageEventOperationResponseBodyData struct {
	// The primary key of the alert handling rule.
	//
	// example:
	//
	// 443496
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteImageEventOperationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageEventOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteImageEventOperationResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DeleteImageEventOperationResponseBodyData) SetId(v int64) *DeleteImageEventOperationResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteImageEventOperationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
