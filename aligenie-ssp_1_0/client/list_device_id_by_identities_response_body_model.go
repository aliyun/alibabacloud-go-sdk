// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceIdByIdentitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListDeviceIdByIdentitiesResponseBody
	GetCode() *int32
	SetMessage(v string) *ListDeviceIdByIdentitiesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDeviceIdByIdentitiesResponseBody
	GetRequestId() *string
	SetResult(v map[string]*ResultValue) *ListDeviceIdByIdentitiesResponseBody
	GetResult() map[string]*ResultValue
}

type ListDeviceIdByIdentitiesResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]*ResultValue `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ListDeviceIdByIdentitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceIdByIdentitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceIdByIdentitiesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListDeviceIdByIdentitiesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDeviceIdByIdentitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceIdByIdentitiesResponseBody) GetResult() map[string]*ResultValue {
	return s.Result
}

func (s *ListDeviceIdByIdentitiesResponseBody) SetCode(v int32) *ListDeviceIdByIdentitiesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDeviceIdByIdentitiesResponseBody) SetMessage(v string) *ListDeviceIdByIdentitiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDeviceIdByIdentitiesResponseBody) SetRequestId(v string) *ListDeviceIdByIdentitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceIdByIdentitiesResponseBody) SetResult(v map[string]*ResultValue) *ListDeviceIdByIdentitiesResponseBody {
	s.Result = v
	return s
}

func (s *ListDeviceIdByIdentitiesResponseBody) Validate() error {
	return dara.Validate(s)
}
