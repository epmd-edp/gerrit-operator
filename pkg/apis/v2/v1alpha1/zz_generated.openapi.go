// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.Gerrit":                        schema_pkg_apis_v2_v1alpha1_Gerrit(ref),
		"github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritReplicationConfig":       schema_pkg_apis_v2_v1alpha1_GerritReplicationConfig(ref),
		"github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritReplicationConfigSpec":   schema_pkg_apis_v2_v1alpha1_GerritReplicationConfigSpec(ref),
		"github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritReplicationConfigStatus": schema_pkg_apis_v2_v1alpha1_GerritReplicationConfigStatus(ref),
		"github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritStatus":                  schema_pkg_apis_v2_v1alpha1_GerritStatus(ref),
	}
}

func schema_pkg_apis_v2_v1alpha1_Gerrit(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Gerrit is the Schema for the gerrits API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritSpec", "github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v2_v1alpha1_GerritReplicationConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GerritReplicationConfig is the Schema for the gerritreplicationconfigs API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritReplicationConfigSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritReplicationConfigStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritReplicationConfigSpec", "github.com/epmd-edp/gerrit-operator/v2/pkg/apis/v2/v1alpha1.GerritReplicationConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_v2_v1alpha1_GerritReplicationConfigSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GerritReplicationConfigSpec defines the desired state of GerritReplicationConfig",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_v2_v1alpha1_GerritReplicationConfigStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GerritReplicationConfigStatus defines the observed state of GerritReplicationConfig",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_v2_v1alpha1_GerritStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GerritStatus defines the observed state of Gerrit",
				Properties: map[string]spec.Schema{
					"available": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"lastTimeUpdated": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "date-time",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"externalUrl": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"externalUrl"},
			},
		},
		Dependencies: []string{},
	}
}
