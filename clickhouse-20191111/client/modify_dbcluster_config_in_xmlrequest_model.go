// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterConfigInXMLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *ModifyDBClusterConfigInXMLRequest
	GetConfig() *string
	SetDBClusterId(v string) *ModifyDBClusterConfigInXMLRequest
	GetDBClusterId() *string
	SetReason(v string) *ModifyDBClusterConfigInXMLRequest
	GetReason() *string
	SetRegionId(v string) *ModifyDBClusterConfigInXMLRequest
	GetRegionId() *string
}

type ModifyDBClusterConfigInXMLRequest struct {
	// The configuration parameters to modify. Call the [DescribeDBClusterConfigInXML](https://help.aliyun.com/document_detail/452210.html) operation to query the current configuration parameters. Then, modify the returned parameters.
	//
	// > When you modify a parameter, you must include all other configuration parameters in the request. If a parameter is not included, it retains its initial value or the modification fails.
	//
	// This parameter is required.
	//
	// example:
	//
	// <?xml version="1.0"?> <yandex>  <logger>    <level>debug</level>    <size>1000M</size>  </logger>  <listen_backlog>64</listen_backlog>  <max_connections>4096</max_connections>  <keep_alive_timeout>300</keep_alive_timeout>  <max_concurrent_queries>100</max_concurrent_queries>  <max_server_memory_usage>0</max_server_memory_usage>  <max_thread_pool_size>10000</max_thread_pool_size>  <max_server_memory_usage_to_ram_ratio>0.9</max_server_memory_usage_to_ram_ratio>  <total_memory_profiler_step>4194304</total_memory_profiler_step>  <total_memory_tracker_sample_probability>0</total_memory_tracker_sample_probability>  <query_log>    <database>system</database>    <table>query_log</table>    <engine>ENGINE = MergeTree PARTITION BY event_date ORDER BY event_time TTL event_time + INTERVAL 15 DAY</engine>    <flush_interval_milliseconds>7500</flush_interval_milliseconds>  </query_log>  <query_thread_log>    <database>system</database>    <table>query_thread_log</table>    <engine>ENGINE = MergeTree PARTITION BY event_date ORDER BY event_time TTL event_time + INTERVAL 15 DAY</engine>    <flush_interval_milliseconds>7500</flush_interval_milliseconds>  </query_thread_log>  <max_table_size_to_drop>0</max_table_size_to_drop>  <max_partition_size_to_drop>0</max_partition_size_to_drop>  <merge_tree>    <max_delay_to_insert>256</max_delay_to_insert>    <max_suspicious_broken_parts>100</max_suspicious_broken_parts>    <zookeeper_session_expiration_check_period>1</zookeeper_session_expiration_check_period>    <max_part_loading_threads>16</max_part_loading_threads>  </merge_tree>  <mark_cache_size>6871947673</mark_cache_size>  <uncompressed_cache_size>3435973836</uncompressed_cache_size></yandex>
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID. Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query the IDs of all clusters in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The reason for the modification.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID. Call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBClusterConfigInXMLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterConfigInXMLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigInXMLRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyDBClusterConfigInXMLRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterConfigInXMLRequest) GetReason() *string {
	return s.Reason
}

func (s *ModifyDBClusterConfigInXMLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterConfigInXMLRequest) SetConfig(v string) *ModifyDBClusterConfigInXMLRequest {
	s.Config = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLRequest) SetDBClusterId(v string) *ModifyDBClusterConfigInXMLRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLRequest) SetReason(v string) *ModifyDBClusterConfigInXMLRequest {
	s.Reason = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLRequest) SetRegionId(v string) *ModifyDBClusterConfigInXMLRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterConfigInXMLRequest) Validate() error {
	return dara.Validate(s)
}
