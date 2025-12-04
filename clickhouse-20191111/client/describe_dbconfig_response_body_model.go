// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeDBConfigResponseBody
	GetConfig() *string
	SetRequestId(v string) *DescribeDBConfigResponseBody
	GetRequestId() *string
}

type DescribeDBConfigResponseBody struct {
	// The configuration information about the cluster.
	//
	// example:
	//
	// <dictionaries><dictionary><name>test</name><source><clickhouse><host>10.37.XX.XX</host><port>9000</port><user>default</user><password></password><db>default</db><table>t_organization</table><where>id=1</where><invalidate_query>SQL_QUERY</invalidate_query></clickhouse></source><lifetime><min>31</min><max>33</max></lifetime><layout><flat></flat></layout><structure><key><attribute><name>field1</name><type>String</type></attribute></key></structure></dictionary></dictionaries>
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16060117-92E1-5F3B-BF42-28B172D0F869
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBConfigResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeDBConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBConfigResponseBody) SetConfig(v string) *DescribeDBConfigResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDBConfigResponseBody) SetRequestId(v string) *DescribeDBConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
