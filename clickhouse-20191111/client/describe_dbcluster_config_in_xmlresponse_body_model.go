// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterConfigInXMLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *DescribeDBClusterConfigInXMLResponseBody
	GetConfig() *string
	SetRequestId(v string) *DescribeDBClusterConfigInXMLResponseBody
	GetRequestId() *string
}

type DescribeDBClusterConfigInXMLResponseBody struct {
	// The values of the configuration parameters.
	//
	// example:
	//
	// <?xml version="1.0"?>
	//
	// <yandex>
	//
	//     <keep_alive_timeout>300</keep_alive_timeout>
	//
	//     <listen_backlog>64</listen_backlog>
	//
	//     <logger>
	//
	//         <level>debug</level>
	//
	//         <size>1000M</size>
	//
	//     </logger>
	//
	//     <mark_cache_size>6871947673</mark_cache_size>
	//
	//     <max_concurrent_queries>100</max_concurrent_queries>
	//
	//     <max_connections>4096</max_connections>
	//
	//     <max_partition_size_to_drop>0</max_partition_size_to_drop>
	//
	//     <max_server_memory_usage>0</max_server_memory_usage>
	//
	// <max_server_memory_usage_to_ram_ratio>0.9</max_server_memory_usage_to_ram_ratio>
	//
	//     <max_table_size_to_drop>0</max_table_size_to_drop>
	//
	//     <max_thread_pool_size>10000</max_thread_pool_size>
	//
	//     <merge_tree>
	//
	//         <max_delay_to_insert>256</max_delay_to_insert>
	//
	//         <max_part_loading_threads>16</max_part_loading_threads>      <zookeeper_session_expiration_check_period>1</zookeeper_session_expiration_check_period>
	//
	//     </merge_tree>
	//
	//     <total_memory_profiler_step>4194304</total_memory_profiler_step>
	//
	// <total_memory_tracker_sample_probability>0</total_memory_tracker_sample_probability>
	//
	//     <uncompressed_cache_size>3435973836</uncompressed_cache_size>
	//
	// </yandex>
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE042911-2B00-134C-9F42-816871BBAFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterConfigInXMLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterConfigInXMLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterConfigInXMLResponseBody) GetConfig() *string {
	return s.Config
}

func (s *DescribeDBClusterConfigInXMLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterConfigInXMLResponseBody) SetConfig(v string) *DescribeDBClusterConfigInXMLResponseBody {
	s.Config = &v
	return s
}

func (s *DescribeDBClusterConfigInXMLResponseBody) SetRequestId(v string) *DescribeDBClusterConfigInXMLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterConfigInXMLResponseBody) Validate() error {
	return dara.Validate(s)
}
