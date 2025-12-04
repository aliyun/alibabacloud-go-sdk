// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSyndbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int64) *DeleteSyndbResponseBody
	GetErrorCode() *int64
	SetErrorMsg(v string) *DeleteSyndbResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *DeleteSyndbResponseBody
	GetRequestId() *string
	SetStatus(v bool) *DeleteSyndbResponseBody
	GetStatus() *bool
}

type DeleteSyndbResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 	- If the value **true*	- is returned for the **Status*	- parameter, the system does not return the ErrorMsg parameter.
	//
	// 	- If the value **false*	- is returned for the **Status*	- parameter, the system returns the deletion failure cause for the ErrorMsg parameter.
	//
	// example:
	//
	// ClickHouse exception, code: 49, host: 100.100.xx.xx, port: 49670; Code: 49, e.displayText() = DB::Exception: Logical error: there is no global context (version 20.8.17.25)n
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2C7393F1-5FD1-5CEE-A2EA-270A2CF99693
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the database used for data synchronization was deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteSyndbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSyndbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSyndbResponseBody) GetErrorCode() *int64 {
	return s.ErrorCode
}

func (s *DeleteSyndbResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteSyndbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSyndbResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *DeleteSyndbResponseBody) SetErrorCode(v int64) *DeleteSyndbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSyndbResponseBody) SetErrorMsg(v string) *DeleteSyndbResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteSyndbResponseBody) SetRequestId(v string) *DeleteSyndbResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSyndbResponseBody) SetStatus(v bool) *DeleteSyndbResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteSyndbResponseBody) Validate() error {
	return dara.Validate(s)
}
