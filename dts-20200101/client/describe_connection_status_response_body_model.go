// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConnectionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationConnectionStatus(v map[string]interface{}) *DescribeConnectionStatusResponseBody
	GetDestinationConnectionStatus() map[string]interface{}
	SetErrCode(v string) *DescribeConnectionStatusResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeConnectionStatusResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DescribeConnectionStatusResponseBody
	GetRequestId() *string
	SetSourceConnectionStatus(v map[string]interface{}) *DescribeConnectionStatusResponseBody
	GetSourceConnectionStatus() map[string]interface{}
	SetSuccess(v string) *DescribeConnectionStatusResponseBody
	GetSuccess() *string
}

type DescribeConnectionStatusResponseBody struct {
	// The connectivity of DTS servers to the destination database.
	//
	// example:
	//
	// {     "connectDetail": [       {         "testName": "PolarDB_o JDBC Connect",         "testSuccess": true       },       {         "testName": "Ping ",         "testSuccess": true       },       {         "testName": "Telnet ",         "testSuccess": true       }     ],     "connectRes": true,     "connectAdvice": ""   }
	DestinationConnectionStatus map[string]interface{} `json:"DestinationConnectionStatus,omitempty" xml:"DestinationConnectionStatus,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0A47C784-70EF-4111-8677-369CAA00****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The connectivity of DTS servers to the source database.
	//
	// example:
	//
	// {     "connectDetail": [       {         "testName": "Oracle JDBC Connect",         "testSuccess": true       },       {         "testName": "Ping ",         "testSuccess": false       },       {         "testName": "Telnet ",         "testSuccess": true       }     ],     "connectRes": true,     "connectAdvice": ""   }
	SourceConnectionStatus map[string]interface{} `json:"SourceConnectionStatus,omitempty" xml:"SourceConnectionStatus,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeConnectionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectionStatusResponseBody) GetDestinationConnectionStatus() map[string]interface{} {
	return s.DestinationConnectionStatus
}

func (s *DescribeConnectionStatusResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeConnectionStatusResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeConnectionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConnectionStatusResponseBody) GetSourceConnectionStatus() map[string]interface{} {
	return s.SourceConnectionStatus
}

func (s *DescribeConnectionStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeConnectionStatusResponseBody) SetDestinationConnectionStatus(v map[string]interface{}) *DescribeConnectionStatusResponseBody {
	s.DestinationConnectionStatus = v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetErrCode(v string) *DescribeConnectionStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetErrMessage(v string) *DescribeConnectionStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetRequestId(v string) *DescribeConnectionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetSourceConnectionStatus(v map[string]interface{}) *DescribeConnectionStatusResponseBody {
	s.SourceConnectionStatus = v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetSuccess(v string) *DescribeConnectionStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
