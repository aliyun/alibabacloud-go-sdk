// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallInPoolTransferConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCallInPoolTransferConfigResponseBody
	GetCode() *string
	SetData(v *QueryCallInPoolTransferConfigResponseBodyData) *QueryCallInPoolTransferConfigResponseBody
	GetData() *QueryCallInPoolTransferConfigResponseBodyData
	SetMessage(v string) *QueryCallInPoolTransferConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCallInPoolTransferConfigResponseBody
	GetRequestId() *string
}

type QueryCallInPoolTransferConfigResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *QueryCallInPoolTransferConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCallInPoolTransferConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInPoolTransferConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCallInPoolTransferConfigResponseBody) GetData() *QueryCallInPoolTransferConfigResponseBodyData {
	return s.Data
}

func (s *QueryCallInPoolTransferConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCallInPoolTransferConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCallInPoolTransferConfigResponseBody) SetCode(v string) *QueryCallInPoolTransferConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBody) SetData(v *QueryCallInPoolTransferConfigResponseBodyData) *QueryCallInPoolTransferConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBody) SetMessage(v string) *QueryCallInPoolTransferConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBody) SetRequestId(v string) *QueryCallInPoolTransferConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCallInPoolTransferConfigResponseBodyData struct {
	// The call mode. Valid values:
	//
	// 	- **roundRobin**
	//
	// 	- **random**
	//
	// example:
	//
	// random
	CalledRouteMode *string `json:"CalledRouteMode,omitempty" xml:"CalledRouteMode,omitempty"`
	// The details of the response parameters.
	Details []*QueryCallInPoolTransferConfigResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The time when the call transfer task was created.
	//
	// example:
	//
	// 1623137002000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timeout period for transferring the call.
	//
	// example:
	//
	// 30
	TransferTimeout *string `json:"TransferTimeout,omitempty" xml:"TransferTimeout,omitempty"`
}

func (s QueryCallInPoolTransferConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInPoolTransferConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) GetCalledRouteMode() *string {
	return s.CalledRouteMode
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) GetDetails() []*QueryCallInPoolTransferConfigResponseBodyDataDetails {
	return s.Details
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) GetTransferTimeout() *string {
	return s.TransferTimeout
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) SetCalledRouteMode(v string) *QueryCallInPoolTransferConfigResponseBodyData {
	s.CalledRouteMode = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) SetDetails(v []*QueryCallInPoolTransferConfigResponseBodyDataDetails) *QueryCallInPoolTransferConfigResponseBodyData {
	s.Details = v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) SetGmtCreate(v int64) *QueryCallInPoolTransferConfigResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) SetTransferTimeout(v string) *QueryCallInPoolTransferConfigResponseBodyData {
	s.TransferTimeout = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCallInPoolTransferConfigResponseBodyDataDetails struct {
	// The number used to transfer the call.
	//
	// example:
	//
	// 400****
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
}

func (s QueryCallInPoolTransferConfigResponseBodyDataDetails) String() string {
	return dara.Prettify(s)
}

func (s QueryCallInPoolTransferConfigResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigResponseBodyDataDetails) GetCalled() *string {
	return s.Called
}

func (s *QueryCallInPoolTransferConfigResponseBodyDataDetails) SetCalled(v string) *QueryCallInPoolTransferConfigResponseBodyDataDetails {
	s.Called = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBodyDataDetails) Validate() error {
	return dara.Validate(s)
}
