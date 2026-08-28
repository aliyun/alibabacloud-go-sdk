// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayErrorAccessLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []map[string]*string) *ListGatewayErrorAccessLogsResponseBody
	GetData() []map[string]*string
	SetRequestId(v string) *ListGatewayErrorAccessLogsResponseBody
	GetRequestId() *string
}

type ListGatewayErrorAccessLogsResponseBody struct {
	Data []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// C9BF49BD-3037-5006-B379-656ECBA6F56C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayErrorAccessLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayErrorAccessLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayErrorAccessLogsResponseBody) GetData() []map[string]*string {
	return s.Data
}

func (s *ListGatewayErrorAccessLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayErrorAccessLogsResponseBody) SetData(v []map[string]*string) *ListGatewayErrorAccessLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayErrorAccessLogsResponseBody) SetRequestId(v string) *ListGatewayErrorAccessLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayErrorAccessLogsResponseBody) Validate() error {
	return dara.Validate(s)
}
