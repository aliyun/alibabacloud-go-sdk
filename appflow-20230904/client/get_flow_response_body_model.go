// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlow(v *GetFlowResponseBodyFlow) *GetFlowResponseBody
	GetFlow() *GetFlowResponseBodyFlow
	SetRequestId(v string) *GetFlowResponseBody
	GetRequestId() *string
}

type GetFlowResponseBody struct {
	// The flow object.
	Flow *GetFlowResponseBodyFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B6E1E38D-011F-5368-ADD8-4DC278254AA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowResponseBody) GetFlow() *GetFlowResponseBodyFlow {
	return s.Flow
}

func (s *GetFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFlowResponseBody) SetFlow(v *GetFlowResponseBodyFlow) *GetFlowResponseBody {
	s.Flow = v
	return s
}

func (s *GetFlowResponseBody) SetRequestId(v string) *GetFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowResponseBody) Validate() error {
	if s.Flow != nil {
		if err := s.Flow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFlowResponseBodyFlow struct {
	// Indicates whether the flow is enabled.
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The flow description.
	//
	// example:
	//
	// 以AI卡片形式发送至钉钉群聊，如果想要支持私聊，请使用同时支持群聊&私聊的模版
	FlowDesc *string `json:"FlowDesc,omitempty" xml:"FlowDesc,omitempty"`
	// The flow ID.
	//
	// example:
	//
	// flow-xxxxxxxx
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The flow name.
	//
	// example:
	//
	// 微信连接流1
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The list of nodes.
	//
	// example:
	//
	// 连接流节点信息
	FlowNodes []*GetFlowResponseBodyFlowFlowNodes `json:"FlowNodes,omitempty" xml:"FlowNodes,omitempty" type:"Repeated"`
	// The flow template content.
	//
	// example:
	//
	// {
	//
	//   "FormatVersion": "appflow-2025-07-01",
	//
	//   "Nodes": [
	//
	//         {}
	//
	//    ]
	//
	// }
	FlowTemplate *string `json:"FlowTemplate,omitempty" xml:"FlowTemplate,omitempty"`
	// The flow version.
	//
	// example:
	//
	// 2
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// The flow version status.
	//
	// example:
	//
	// 1
	FlowVersionStatus *string `json:"FlowVersionStatus,omitempty" xml:"FlowVersionStatus,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-07-30T02:13:22Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2025-11-13T02:11:56Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// A list of tags.
	Tags []*GetFlowResponseBodyFlowTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetFlowResponseBodyFlow) String() string {
	return dara.Prettify(s)
}

func (s GetFlowResponseBodyFlow) GoString() string {
	return s.String()
}

func (s *GetFlowResponseBodyFlow) GetEnabled() *string {
	return s.Enabled
}

func (s *GetFlowResponseBodyFlow) GetFlowDesc() *string {
	return s.FlowDesc
}

func (s *GetFlowResponseBodyFlow) GetFlowId() *string {
	return s.FlowId
}

func (s *GetFlowResponseBodyFlow) GetFlowName() *string {
	return s.FlowName
}

func (s *GetFlowResponseBodyFlow) GetFlowNodes() []*GetFlowResponseBodyFlowFlowNodes {
	return s.FlowNodes
}

func (s *GetFlowResponseBodyFlow) GetFlowTemplate() *string {
	return s.FlowTemplate
}

func (s *GetFlowResponseBodyFlow) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *GetFlowResponseBodyFlow) GetFlowVersionStatus() *string {
	return s.FlowVersionStatus
}

func (s *GetFlowResponseBodyFlow) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetFlowResponseBodyFlow) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetFlowResponseBodyFlow) GetTags() []*GetFlowResponseBodyFlowTags {
	return s.Tags
}

func (s *GetFlowResponseBodyFlow) SetEnabled(v string) *GetFlowResponseBodyFlow {
	s.Enabled = &v
	return s
}

func (s *GetFlowResponseBodyFlow) SetFlowDesc(v string) *GetFlowResponseBodyFlow {
	s.FlowDesc = &v
	return s
}

func (s *GetFlowResponseBodyFlow) SetFlowId(v string) *GetFlowResponseBodyFlow {
	s.FlowId = &v
	return s
}

func (s *GetFlowResponseBodyFlow) SetFlowName(v string) *GetFlowResponseBodyFlow {
	s.FlowName = &v
	return s
}

func (s *GetFlowResponseBodyFlow) SetFlowNodes(v []*GetFlowResponseBodyFlowFlowNodes) *GetFlowResponseBodyFlow {
	s.FlowNodes = v
	return s
}

func (s *GetFlowResponseBodyFlow) SetFlowTemplate(v string) *GetFlowResponseBodyFlow {
	s.FlowTemplate = &v
	return s
}

func (s *GetFlowResponseBodyFlow) SetFlowVersion(v string) *GetFlowResponseBodyFlow {
	s.FlowVersion = &v
	return s
}

func (s *GetFlowResponseBodyFlow) SetFlowVersionStatus(v string) *GetFlowResponseBodyFlow {
	s.FlowVersionStatus = &v
	return s
}

func (s *GetFlowResponseBodyFlow) SetGmtCreate(v string) *GetFlowResponseBodyFlow {
	s.GmtCreate = &v
	return s
}

func (s *GetFlowResponseBodyFlow) SetGmtModified(v string) *GetFlowResponseBodyFlow {
	s.GmtModified = &v
	return s
}

func (s *GetFlowResponseBodyFlow) SetTags(v []*GetFlowResponseBodyFlowTags) *GetFlowResponseBodyFlow {
	s.Tags = v
	return s
}

func (s *GetFlowResponseBodyFlow) Validate() error {
	if s.FlowNodes != nil {
		for _, item := range s.FlowNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFlowResponseBodyFlowFlowNodes struct {
	// The authentication credentials of the node.
	//
	// example:
	//
	// {\\"authconfigId\\":\\"uac-xxxxxxxxx\\"}
	AuthMetadata *string `json:"AuthMetadata,omitempty" xml:"AuthMetadata,omitempty"`
	// The connector ID.
	//
	// example:
	//
	// connector-xxx24b139c62
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The connector version.
	//
	// example:
	//
	// 2
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
	// The flow ID.
	//
	// example:
	//
	// flow-856cb84b309747e48b43
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// The flow version.
	//
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// The metadata of the node.
	//
	// example:
	//
	// {}
	InputSchema *string `json:"InputSchema,omitempty" xml:"InputSchema,omitempty"`
	// The node ID.
	//
	// example:
	//
	// fn-xxxxxxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node key.
	//
	// example:
	//
	// Node1
	NodeKey *string `json:"NodeKey,omitempty" xml:"NodeKey,omitempty"`
	// The node name.
	//
	// example:
	//
	// NotifyMessage_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The node type.
	//
	// example:
	//
	// Trigger
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The ID of the previous node.
	//
	// example:
	//
	// fn-xxxxx,fn-yyyyyy
	PrevNodeId *string `json:"PrevNodeId,omitempty" xml:"PrevNodeId,omitempty"`
	// The ID of the trigger or action.
	//
	// example:
	//
	// trigger-xxxxxxxxxx
	RefId *string `json:"RefId,omitempty" xml:"RefId,omitempty"`
	// The trigger or action version.
	//
	// example:
	//
	// 1
	RefVersion *string `json:"RefVersion,omitempty" xml:"RefVersion,omitempty"`
	// The webhook URL.
	//
	// example:
	//
	// https://{uid}.computenest.aliyun.com/webhook/xxxxxxxx
	WebhookUrl *string `json:"WebhookUrl,omitempty" xml:"WebhookUrl,omitempty"`
}

func (s GetFlowResponseBodyFlowFlowNodes) String() string {
	return dara.Prettify(s)
}

func (s GetFlowResponseBodyFlowFlowNodes) GoString() string {
	return s.String()
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetAuthMetadata() *string {
	return s.AuthMetadata
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetFlowId() *string {
	return s.FlowId
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetInputSchema() *string {
	return s.InputSchema
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetNodeKey() *string {
	return s.NodeKey
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetNodeType() *string {
	return s.NodeType
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetPrevNodeId() *string {
	return s.PrevNodeId
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetRefId() *string {
	return s.RefId
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetRefVersion() *string {
	return s.RefVersion
}

func (s *GetFlowResponseBodyFlowFlowNodes) GetWebhookUrl() *string {
	return s.WebhookUrl
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetAuthMetadata(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.AuthMetadata = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetConnectorId(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.ConnectorId = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetConnectorVersion(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.ConnectorVersion = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetFlowId(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.FlowId = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetFlowVersion(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.FlowVersion = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetInputSchema(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.InputSchema = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetNodeId(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.NodeId = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetNodeKey(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.NodeKey = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetNodeName(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.NodeName = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetNodeType(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.NodeType = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetPrevNodeId(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.PrevNodeId = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetRefId(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.RefId = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetRefVersion(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.RefVersion = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) SetWebhookUrl(v string) *GetFlowResponseBodyFlowFlowNodes {
	s.WebhookUrl = &v
	return s
}

func (s *GetFlowResponseBodyFlowFlowNodes) Validate() error {
	return dara.Validate(s)
}

type GetFlowResponseBodyFlowTags struct {
	// The tag key. The value can be up to 64 characters in length.
	//
	// example:
	//
	// Environment
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// pre
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetFlowResponseBodyFlowTags) String() string {
	return dara.Prettify(s)
}

func (s GetFlowResponseBodyFlowTags) GoString() string {
	return s.String()
}

func (s *GetFlowResponseBodyFlowTags) GetKey() *string {
	return s.Key
}

func (s *GetFlowResponseBodyFlowTags) GetValue() *string {
	return s.Value
}

func (s *GetFlowResponseBodyFlowTags) SetKey(v string) *GetFlowResponseBodyFlowTags {
	s.Key = &v
	return s
}

func (s *GetFlowResponseBodyFlowTags) SetValue(v string) *GetFlowResponseBodyFlowTags {
	s.Value = &v
	return s
}

func (s *GetFlowResponseBodyFlowTags) Validate() error {
	return dara.Validate(s)
}
