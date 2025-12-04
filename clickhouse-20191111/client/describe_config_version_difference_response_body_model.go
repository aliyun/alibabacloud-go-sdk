// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigVersionDifferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNewConfigXML(v string) *DescribeConfigVersionDifferenceResponseBody
	GetNewConfigXML() *string
	SetOldConfigXML(v string) *DescribeConfigVersionDifferenceResponseBody
	GetOldConfigXML() *string
	SetRequestId(v string) *DescribeConfigVersionDifferenceResponseBody
	GetRequestId() *string
}

type DescribeConfigVersionDifferenceResponseBody struct {
	// The values of the configuration parameters after the values of the configuration parameters are changed.
	//
	// example:
	//
	// "<?xml version="1.0"?>
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
	// </yandex>"
	NewConfigXML *string `json:"NewConfigXML,omitempty" xml:"NewConfigXML,omitempty"`
	// The values of the configuration parameters before the values of the configuration parameters are changed.
	//
	// example:
	//
	// "<?xml version="1.0"?>
	//
	// <yandex>
	//
	//     <keep_alive_timeout>300</keep_alive_timeout>
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
	//     <max_concurrent_queries>150</max_concurrent_queries>
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
	// </yandex>"
	OldConfigXML *string `json:"OldConfigXML,omitempty" xml:"OldConfigXML,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConfigVersionDifferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigVersionDifferenceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigVersionDifferenceResponseBody) GetNewConfigXML() *string {
	return s.NewConfigXML
}

func (s *DescribeConfigVersionDifferenceResponseBody) GetOldConfigXML() *string {
	return s.OldConfigXML
}

func (s *DescribeConfigVersionDifferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConfigVersionDifferenceResponseBody) SetNewConfigXML(v string) *DescribeConfigVersionDifferenceResponseBody {
	s.NewConfigXML = &v
	return s
}

func (s *DescribeConfigVersionDifferenceResponseBody) SetOldConfigXML(v string) *DescribeConfigVersionDifferenceResponseBody {
	s.OldConfigXML = &v
	return s
}

func (s *DescribeConfigVersionDifferenceResponseBody) SetRequestId(v string) *DescribeConfigVersionDifferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigVersionDifferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
