// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWildcardByTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteWildcardByTopicResponseBody
	GetCode() *int32
	SetData(v *DeleteWildcardByTopicResponseBodyData) *DeleteWildcardByTopicResponseBody
	GetData() *DeleteWildcardByTopicResponseBodyData
	SetRequestId(v string) *DeleteWildcardByTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWildcardByTopicResponseBody
	GetSuccess() *bool
}

type DeleteWildcardByTopicResponseBody struct {
	// The response code. Valid values:
	//
	// - 200: success.
	//
	// - Other values: error codes. For more information, see error codes.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The rule information dataset.
	Data *DeleteWildcardByTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82B9E503-F4A1-4F30-976F-C6999FF9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWildcardByTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWildcardByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWildcardByTopicResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteWildcardByTopicResponseBody) GetData() *DeleteWildcardByTopicResponseBodyData {
	return s.Data
}

func (s *DeleteWildcardByTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWildcardByTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWildcardByTopicResponseBody) SetCode(v int32) *DeleteWildcardByTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWildcardByTopicResponseBody) SetData(v *DeleteWildcardByTopicResponseBodyData) *DeleteWildcardByTopicResponseBody {
	s.Data = v
	return s
}

func (s *DeleteWildcardByTopicResponseBody) SetRequestId(v string) *DeleteWildcardByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWildcardByTopicResponseBody) SetSuccess(v bool) *DeleteWildcardByTopicResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWildcardByTopicResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteWildcardByTopicResponseBodyData struct {
	// Indicates whether the deletion was successful. Valid values:
	//
	// - **true**: The deletion was successful.
	//
	// - **false**: The deletion failed.
	//
	// example:
	//
	// True
	Deleted *string `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
}

func (s DeleteWildcardByTopicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteWildcardByTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteWildcardByTopicResponseBodyData) GetDeleted() *string {
	return s.Deleted
}

func (s *DeleteWildcardByTopicResponseBodyData) SetDeleted(v string) *DeleteWildcardByTopicResponseBodyData {
	s.Deleted = &v
	return s
}

func (s *DeleteWildcardByTopicResponseBodyData) Validate() error {
	return dara.Validate(s)
}
