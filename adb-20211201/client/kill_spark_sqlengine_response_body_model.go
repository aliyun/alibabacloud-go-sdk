// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSparkSQLEngineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *KillSparkSQLEngineResponseBody
	GetData() *bool
	SetRequestId(v string) *KillSparkSQLEngineResponseBody
	GetRequestId() *string
}

type KillSparkSQLEngineResponseBody struct {
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillSparkSQLEngineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KillSparkSQLEngineResponseBody) GoString() string {
	return s.String()
}

func (s *KillSparkSQLEngineResponseBody) GetData() *bool {
	return s.Data
}

func (s *KillSparkSQLEngineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KillSparkSQLEngineResponseBody) SetData(v bool) *KillSparkSQLEngineResponseBody {
	s.Data = &v
	return s
}

func (s *KillSparkSQLEngineResponseBody) SetRequestId(v string) *KillSparkSQLEngineResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillSparkSQLEngineResponseBody) Validate() error {
	return dara.Validate(s)
}
