// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkCodeLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLog(v string) *DescribeSparkCodeLogResponseBody
	GetLog() *string
	SetMessage(v string) *DescribeSparkCodeLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSparkCodeLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSparkCodeLogResponseBody
	GetSuccess() *bool
}

type DescribeSparkCodeLogResponseBody struct {
	// The content of the log.
	//
	// example:
	//
	// >>>>>>>> stdout:n++++++++++++++++++executing sql: MSCK REPAIR TABLE  `footprint_ethereum`.`dwd_eth_eth_txr_v2_di` n++n
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// The returned message.
	//
	// 	- If the request was successful, **Success*	- is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CD65640-9963-5D60-929C-118F2C84070E
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

func (s DescribeSparkCodeLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkCodeLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeLogResponseBody) GetLog() *string {
	return s.Log
}

func (s *DescribeSparkCodeLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSparkCodeLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSparkCodeLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSparkCodeLogResponseBody) SetLog(v string) *DescribeSparkCodeLogResponseBody {
	s.Log = &v
	return s
}

func (s *DescribeSparkCodeLogResponseBody) SetMessage(v string) *DescribeSparkCodeLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSparkCodeLogResponseBody) SetRequestId(v string) *DescribeSparkCodeLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSparkCodeLogResponseBody) SetSuccess(v bool) *DescribeSparkCodeLogResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSparkCodeLogResponseBody) Validate() error {
	return dara.Validate(s)
}
