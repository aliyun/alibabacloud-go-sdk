// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadSampleDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *LoadSampleDataSetResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *LoadSampleDataSetResponseBody
	GetRequestId() *string
}

type LoadSampleDataSetResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-2ze0z517o1mgp66a
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FA31BE84-ABE8-554A-A769-5F860C34EE10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LoadSampleDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoadSampleDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *LoadSampleDataSetResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *LoadSampleDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoadSampleDataSetResponseBody) SetDBClusterId(v string) *LoadSampleDataSetResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *LoadSampleDataSetResponseBody) SetRequestId(v string) *LoadSampleDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoadSampleDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}
