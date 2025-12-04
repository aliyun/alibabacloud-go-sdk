// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckModifyConfigNeedRestartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CheckModifyConfigNeedRestartRequest
	GetConfig() *string
	SetDBClusterId(v string) *CheckModifyConfigNeedRestartRequest
	GetDBClusterId() *string
}

type CheckModifyConfigNeedRestartRequest struct {
	// The configuration parameters whose settings are modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// <?xml version="1.0"?>
	//
	// <yandex>
	//
	//     <keep_alive_timeout>400</keep_alive_timeout>
	//
	//     <listen_backlog>4096</listen_backlog>
	//
	//     <logger>
	//
	//         <level>debug</level>
	//
	//         <size>1000M</size>
	//
	//     </logger>
	//
	//     <mark_cache_size>5368709120</mark_cache_size>
	//
	//     <max_concurrent_queries>201</max_concurrent_queries>
	//
	//     <max_connections>4096</max_connections>
	//
	//     <max_partition_size_to_drop>0</max_partition_size_to_drop>
	//
	//     <max_table_size_to_drop>0</max_table_size_to_drop>
	//
	//     <merge_tree>
	//
	//         <max_delay_to_insert>256</max_delay_to_insert>
	//
	//         <max_part_loading_threads>auto</max_part_loading_threads>
	//
	//         <max_suspicious_broken_parts>100</max_suspicious_broken_parts>
	//
	//         <zookeeper_session_expiration_check_period>1</zookeeper_session_expiration_check_period>
	//
	//     </merge_tree>
	//
	//     <uncompressed_cache_size>1717986918</uncompressed_cache_size>
	//
	// </yandex>
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query information about all the clusters that are deployed in a specific region. The information includes the cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1tm8zf130ew****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s CheckModifyConfigNeedRestartRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckModifyConfigNeedRestartRequest) GoString() string {
	return s.String()
}

func (s *CheckModifyConfigNeedRestartRequest) GetConfig() *string {
	return s.Config
}

func (s *CheckModifyConfigNeedRestartRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CheckModifyConfigNeedRestartRequest) SetConfig(v string) *CheckModifyConfigNeedRestartRequest {
	s.Config = &v
	return s
}

func (s *CheckModifyConfigNeedRestartRequest) SetDBClusterId(v string) *CheckModifyConfigNeedRestartRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckModifyConfigNeedRestartRequest) Validate() error {
	return dara.Validate(s)
}
