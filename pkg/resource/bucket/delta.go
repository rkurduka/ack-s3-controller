// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package bucket

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customPreCompare(a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.ACL, b.ko.Spec.ACL) {
		delta.Add("Spec.ACL", a.ko.Spec.ACL, b.ko.Spec.ACL)
	} else if a.ko.Spec.ACL != nil && b.ko.Spec.ACL != nil {
		if *a.ko.Spec.ACL != *b.ko.Spec.ACL {
			delta.Add("Spec.ACL", a.ko.Spec.ACL, b.ko.Spec.ACL)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Accelerate, b.ko.Spec.Accelerate) {
		delta.Add("Spec.Accelerate", a.ko.Spec.Accelerate, b.ko.Spec.Accelerate)
	} else if a.ko.Spec.Accelerate != nil && b.ko.Spec.Accelerate != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.Accelerate.Status, b.ko.Spec.Accelerate.Status) {
			delta.Add("Spec.Accelerate.Status", a.ko.Spec.Accelerate.Status, b.ko.Spec.Accelerate.Status)
		} else if a.ko.Spec.Accelerate.Status != nil && b.ko.Spec.Accelerate.Status != nil {
			if *a.ko.Spec.Accelerate.Status != *b.ko.Spec.Accelerate.Status {
				delta.Add("Spec.Accelerate.Status", a.ko.Spec.Accelerate.Status, b.ko.Spec.Accelerate.Status)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CORS, b.ko.Spec.CORS) {
		delta.Add("Spec.CORS", a.ko.Spec.CORS, b.ko.Spec.CORS)
	} else if a.ko.Spec.CORS != nil && b.ko.Spec.CORS != nil {
		if !reflect.DeepEqual(a.ko.Spec.CORS.CORSRules, b.ko.Spec.CORS.CORSRules) {
			delta.Add("Spec.CORS.CORSRules", a.ko.Spec.CORS.CORSRules, b.ko.Spec.CORS.CORSRules)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CreateBucketConfiguration, b.ko.Spec.CreateBucketConfiguration) {
		delta.Add("Spec.CreateBucketConfiguration", a.ko.Spec.CreateBucketConfiguration, b.ko.Spec.CreateBucketConfiguration)
	} else if a.ko.Spec.CreateBucketConfiguration != nil && b.ko.Spec.CreateBucketConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.CreateBucketConfiguration.LocationConstraint, b.ko.Spec.CreateBucketConfiguration.LocationConstraint) {
			delta.Add("Spec.CreateBucketConfiguration.LocationConstraint", a.ko.Spec.CreateBucketConfiguration.LocationConstraint, b.ko.Spec.CreateBucketConfiguration.LocationConstraint)
		} else if a.ko.Spec.CreateBucketConfiguration.LocationConstraint != nil && b.ko.Spec.CreateBucketConfiguration.LocationConstraint != nil {
			if *a.ko.Spec.CreateBucketConfiguration.LocationConstraint != *b.ko.Spec.CreateBucketConfiguration.LocationConstraint {
				delta.Add("Spec.CreateBucketConfiguration.LocationConstraint", a.ko.Spec.CreateBucketConfiguration.LocationConstraint, b.ko.Spec.CreateBucketConfiguration.LocationConstraint)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Encryption, b.ko.Spec.Encryption) {
		delta.Add("Spec.Encryption", a.ko.Spec.Encryption, b.ko.Spec.Encryption)
	} else if a.ko.Spec.Encryption != nil && b.ko.Spec.Encryption != nil {
		if !reflect.DeepEqual(a.ko.Spec.Encryption.Rules, b.ko.Spec.Encryption.Rules) {
			delta.Add("Spec.Encryption.Rules", a.ko.Spec.Encryption.Rules, b.ko.Spec.Encryption.Rules)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.GrantFullControl, b.ko.Spec.GrantFullControl) {
		delta.Add("Spec.GrantFullControl", a.ko.Spec.GrantFullControl, b.ko.Spec.GrantFullControl)
	} else if a.ko.Spec.GrantFullControl != nil && b.ko.Spec.GrantFullControl != nil {
		if *a.ko.Spec.GrantFullControl != *b.ko.Spec.GrantFullControl {
			delta.Add("Spec.GrantFullControl", a.ko.Spec.GrantFullControl, b.ko.Spec.GrantFullControl)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.GrantRead, b.ko.Spec.GrantRead) {
		delta.Add("Spec.GrantRead", a.ko.Spec.GrantRead, b.ko.Spec.GrantRead)
	} else if a.ko.Spec.GrantRead != nil && b.ko.Spec.GrantRead != nil {
		if *a.ko.Spec.GrantRead != *b.ko.Spec.GrantRead {
			delta.Add("Spec.GrantRead", a.ko.Spec.GrantRead, b.ko.Spec.GrantRead)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.GrantReadACP, b.ko.Spec.GrantReadACP) {
		delta.Add("Spec.GrantReadACP", a.ko.Spec.GrantReadACP, b.ko.Spec.GrantReadACP)
	} else if a.ko.Spec.GrantReadACP != nil && b.ko.Spec.GrantReadACP != nil {
		if *a.ko.Spec.GrantReadACP != *b.ko.Spec.GrantReadACP {
			delta.Add("Spec.GrantReadACP", a.ko.Spec.GrantReadACP, b.ko.Spec.GrantReadACP)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.GrantWrite, b.ko.Spec.GrantWrite) {
		delta.Add("Spec.GrantWrite", a.ko.Spec.GrantWrite, b.ko.Spec.GrantWrite)
	} else if a.ko.Spec.GrantWrite != nil && b.ko.Spec.GrantWrite != nil {
		if *a.ko.Spec.GrantWrite != *b.ko.Spec.GrantWrite {
			delta.Add("Spec.GrantWrite", a.ko.Spec.GrantWrite, b.ko.Spec.GrantWrite)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.GrantWriteACP, b.ko.Spec.GrantWriteACP) {
		delta.Add("Spec.GrantWriteACP", a.ko.Spec.GrantWriteACP, b.ko.Spec.GrantWriteACP)
	} else if a.ko.Spec.GrantWriteACP != nil && b.ko.Spec.GrantWriteACP != nil {
		if *a.ko.Spec.GrantWriteACP != *b.ko.Spec.GrantWriteACP {
			delta.Add("Spec.GrantWriteACP", a.ko.Spec.GrantWriteACP, b.ko.Spec.GrantWriteACP)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Lifecycle, b.ko.Spec.Lifecycle) {
		delta.Add("Spec.Lifecycle", a.ko.Spec.Lifecycle, b.ko.Spec.Lifecycle)
	} else if a.ko.Spec.Lifecycle != nil && b.ko.Spec.Lifecycle != nil {
		if !reflect.DeepEqual(a.ko.Spec.Lifecycle.Rules, b.ko.Spec.Lifecycle.Rules) {
			delta.Add("Spec.Lifecycle.Rules", a.ko.Spec.Lifecycle.Rules, b.ko.Spec.Lifecycle.Rules)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Logging, b.ko.Spec.Logging) {
		delta.Add("Spec.Logging", a.ko.Spec.Logging, b.ko.Spec.Logging)
	} else if a.ko.Spec.Logging != nil && b.ko.Spec.Logging != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.Logging.LoggingEnabled, b.ko.Spec.Logging.LoggingEnabled) {
			delta.Add("Spec.Logging.LoggingEnabled", a.ko.Spec.Logging.LoggingEnabled, b.ko.Spec.Logging.LoggingEnabled)
		} else if a.ko.Spec.Logging.LoggingEnabled != nil && b.ko.Spec.Logging.LoggingEnabled != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Logging.LoggingEnabled.TargetBucket, b.ko.Spec.Logging.LoggingEnabled.TargetBucket) {
				delta.Add("Spec.Logging.LoggingEnabled.TargetBucket", a.ko.Spec.Logging.LoggingEnabled.TargetBucket, b.ko.Spec.Logging.LoggingEnabled.TargetBucket)
			} else if a.ko.Spec.Logging.LoggingEnabled.TargetBucket != nil && b.ko.Spec.Logging.LoggingEnabled.TargetBucket != nil {
				if *a.ko.Spec.Logging.LoggingEnabled.TargetBucket != *b.ko.Spec.Logging.LoggingEnabled.TargetBucket {
					delta.Add("Spec.Logging.LoggingEnabled.TargetBucket", a.ko.Spec.Logging.LoggingEnabled.TargetBucket, b.ko.Spec.Logging.LoggingEnabled.TargetBucket)
				}
			}
			if !reflect.DeepEqual(a.ko.Spec.Logging.LoggingEnabled.TargetGrants, b.ko.Spec.Logging.LoggingEnabled.TargetGrants) {
				delta.Add("Spec.Logging.LoggingEnabled.TargetGrants", a.ko.Spec.Logging.LoggingEnabled.TargetGrants, b.ko.Spec.Logging.LoggingEnabled.TargetGrants)
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Logging.LoggingEnabled.TargetPrefix, b.ko.Spec.Logging.LoggingEnabled.TargetPrefix) {
				delta.Add("Spec.Logging.LoggingEnabled.TargetPrefix", a.ko.Spec.Logging.LoggingEnabled.TargetPrefix, b.ko.Spec.Logging.LoggingEnabled.TargetPrefix)
			} else if a.ko.Spec.Logging.LoggingEnabled.TargetPrefix != nil && b.ko.Spec.Logging.LoggingEnabled.TargetPrefix != nil {
				if *a.ko.Spec.Logging.LoggingEnabled.TargetPrefix != *b.ko.Spec.Logging.LoggingEnabled.TargetPrefix {
					delta.Add("Spec.Logging.LoggingEnabled.TargetPrefix", a.ko.Spec.Logging.LoggingEnabled.TargetPrefix, b.ko.Spec.Logging.LoggingEnabled.TargetPrefix)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Notification, b.ko.Spec.Notification) {
		delta.Add("Spec.Notification", a.ko.Spec.Notification, b.ko.Spec.Notification)
	} else if a.ko.Spec.Notification != nil && b.ko.Spec.Notification != nil {
		if !reflect.DeepEqual(a.ko.Spec.Notification.LambdaFunctionConfigurations, b.ko.Spec.Notification.LambdaFunctionConfigurations) {
			delta.Add("Spec.Notification.LambdaFunctionConfigurations", a.ko.Spec.Notification.LambdaFunctionConfigurations, b.ko.Spec.Notification.LambdaFunctionConfigurations)
		}
		if !reflect.DeepEqual(a.ko.Spec.Notification.QueueConfigurations, b.ko.Spec.Notification.QueueConfigurations) {
			delta.Add("Spec.Notification.QueueConfigurations", a.ko.Spec.Notification.QueueConfigurations, b.ko.Spec.Notification.QueueConfigurations)
		}
		if !reflect.DeepEqual(a.ko.Spec.Notification.TopicConfigurations, b.ko.Spec.Notification.TopicConfigurations) {
			delta.Add("Spec.Notification.TopicConfigurations", a.ko.Spec.Notification.TopicConfigurations, b.ko.Spec.Notification.TopicConfigurations)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ObjectLockEnabledForBucket, b.ko.Spec.ObjectLockEnabledForBucket) {
		delta.Add("Spec.ObjectLockEnabledForBucket", a.ko.Spec.ObjectLockEnabledForBucket, b.ko.Spec.ObjectLockEnabledForBucket)
	} else if a.ko.Spec.ObjectLockEnabledForBucket != nil && b.ko.Spec.ObjectLockEnabledForBucket != nil {
		if *a.ko.Spec.ObjectLockEnabledForBucket != *b.ko.Spec.ObjectLockEnabledForBucket {
			delta.Add("Spec.ObjectLockEnabledForBucket", a.ko.Spec.ObjectLockEnabledForBucket, b.ko.Spec.ObjectLockEnabledForBucket)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.OwnershipControls, b.ko.Spec.OwnershipControls) {
		delta.Add("Spec.OwnershipControls", a.ko.Spec.OwnershipControls, b.ko.Spec.OwnershipControls)
	} else if a.ko.Spec.OwnershipControls != nil && b.ko.Spec.OwnershipControls != nil {
		if !reflect.DeepEqual(a.ko.Spec.OwnershipControls.Rules, b.ko.Spec.OwnershipControls.Rules) {
			delta.Add("Spec.OwnershipControls.Rules", a.ko.Spec.OwnershipControls.Rules, b.ko.Spec.OwnershipControls.Rules)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Policy, b.ko.Spec.Policy) {
		delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
	} else if a.ko.Spec.Policy != nil && b.ko.Spec.Policy != nil {
		if *a.ko.Spec.Policy != *b.ko.Spec.Policy {
			delta.Add("Spec.Policy", a.ko.Spec.Policy, b.ko.Spec.Policy)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PublicAccessBlock, b.ko.Spec.PublicAccessBlock) {
		delta.Add("Spec.PublicAccessBlock", a.ko.Spec.PublicAccessBlock, b.ko.Spec.PublicAccessBlock)
	} else if a.ko.Spec.PublicAccessBlock != nil && b.ko.Spec.PublicAccessBlock != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.PublicAccessBlock.BlockPublicACLs, b.ko.Spec.PublicAccessBlock.BlockPublicACLs) {
			delta.Add("Spec.PublicAccessBlock.BlockPublicACLs", a.ko.Spec.PublicAccessBlock.BlockPublicACLs, b.ko.Spec.PublicAccessBlock.BlockPublicACLs)
		} else if a.ko.Spec.PublicAccessBlock.BlockPublicACLs != nil && b.ko.Spec.PublicAccessBlock.BlockPublicACLs != nil {
			if *a.ko.Spec.PublicAccessBlock.BlockPublicACLs != *b.ko.Spec.PublicAccessBlock.BlockPublicACLs {
				delta.Add("Spec.PublicAccessBlock.BlockPublicACLs", a.ko.Spec.PublicAccessBlock.BlockPublicACLs, b.ko.Spec.PublicAccessBlock.BlockPublicACLs)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PublicAccessBlock.BlockPublicPolicy, b.ko.Spec.PublicAccessBlock.BlockPublicPolicy) {
			delta.Add("Spec.PublicAccessBlock.BlockPublicPolicy", a.ko.Spec.PublicAccessBlock.BlockPublicPolicy, b.ko.Spec.PublicAccessBlock.BlockPublicPolicy)
		} else if a.ko.Spec.PublicAccessBlock.BlockPublicPolicy != nil && b.ko.Spec.PublicAccessBlock.BlockPublicPolicy != nil {
			if *a.ko.Spec.PublicAccessBlock.BlockPublicPolicy != *b.ko.Spec.PublicAccessBlock.BlockPublicPolicy {
				delta.Add("Spec.PublicAccessBlock.BlockPublicPolicy", a.ko.Spec.PublicAccessBlock.BlockPublicPolicy, b.ko.Spec.PublicAccessBlock.BlockPublicPolicy)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PublicAccessBlock.IgnorePublicACLs, b.ko.Spec.PublicAccessBlock.IgnorePublicACLs) {
			delta.Add("Spec.PublicAccessBlock.IgnorePublicACLs", a.ko.Spec.PublicAccessBlock.IgnorePublicACLs, b.ko.Spec.PublicAccessBlock.IgnorePublicACLs)
		} else if a.ko.Spec.PublicAccessBlock.IgnorePublicACLs != nil && b.ko.Spec.PublicAccessBlock.IgnorePublicACLs != nil {
			if *a.ko.Spec.PublicAccessBlock.IgnorePublicACLs != *b.ko.Spec.PublicAccessBlock.IgnorePublicACLs {
				delta.Add("Spec.PublicAccessBlock.IgnorePublicACLs", a.ko.Spec.PublicAccessBlock.IgnorePublicACLs, b.ko.Spec.PublicAccessBlock.IgnorePublicACLs)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.PublicAccessBlock.RestrictPublicBuckets, b.ko.Spec.PublicAccessBlock.RestrictPublicBuckets) {
			delta.Add("Spec.PublicAccessBlock.RestrictPublicBuckets", a.ko.Spec.PublicAccessBlock.RestrictPublicBuckets, b.ko.Spec.PublicAccessBlock.RestrictPublicBuckets)
		} else if a.ko.Spec.PublicAccessBlock.RestrictPublicBuckets != nil && b.ko.Spec.PublicAccessBlock.RestrictPublicBuckets != nil {
			if *a.ko.Spec.PublicAccessBlock.RestrictPublicBuckets != *b.ko.Spec.PublicAccessBlock.RestrictPublicBuckets {
				delta.Add("Spec.PublicAccessBlock.RestrictPublicBuckets", a.ko.Spec.PublicAccessBlock.RestrictPublicBuckets, b.ko.Spec.PublicAccessBlock.RestrictPublicBuckets)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Replication, b.ko.Spec.Replication) {
		delta.Add("Spec.Replication", a.ko.Spec.Replication, b.ko.Spec.Replication)
	} else if a.ko.Spec.Replication != nil && b.ko.Spec.Replication != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.Replication.Role, b.ko.Spec.Replication.Role) {
			delta.Add("Spec.Replication.Role", a.ko.Spec.Replication.Role, b.ko.Spec.Replication.Role)
		} else if a.ko.Spec.Replication.Role != nil && b.ko.Spec.Replication.Role != nil {
			if *a.ko.Spec.Replication.Role != *b.ko.Spec.Replication.Role {
				delta.Add("Spec.Replication.Role", a.ko.Spec.Replication.Role, b.ko.Spec.Replication.Role)
			}
		}
		if !reflect.DeepEqual(a.ko.Spec.Replication.Rules, b.ko.Spec.Replication.Rules) {
			delta.Add("Spec.Replication.Rules", a.ko.Spec.Replication.Rules, b.ko.Spec.Replication.Rules)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RequestPayment, b.ko.Spec.RequestPayment) {
		delta.Add("Spec.RequestPayment", a.ko.Spec.RequestPayment, b.ko.Spec.RequestPayment)
	} else if a.ko.Spec.RequestPayment != nil && b.ko.Spec.RequestPayment != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.RequestPayment.Payer, b.ko.Spec.RequestPayment.Payer) {
			delta.Add("Spec.RequestPayment.Payer", a.ko.Spec.RequestPayment.Payer, b.ko.Spec.RequestPayment.Payer)
		} else if a.ko.Spec.RequestPayment.Payer != nil && b.ko.Spec.RequestPayment.Payer != nil {
			if *a.ko.Spec.RequestPayment.Payer != *b.ko.Spec.RequestPayment.Payer {
				delta.Add("Spec.RequestPayment.Payer", a.ko.Spec.RequestPayment.Payer, b.ko.Spec.RequestPayment.Payer)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Tagging, b.ko.Spec.Tagging) {
		delta.Add("Spec.Tagging", a.ko.Spec.Tagging, b.ko.Spec.Tagging)
	} else if a.ko.Spec.Tagging != nil && b.ko.Spec.Tagging != nil {
		if !reflect.DeepEqual(a.ko.Spec.Tagging.TagSet, b.ko.Spec.Tagging.TagSet) {
			delta.Add("Spec.Tagging.TagSet", a.ko.Spec.Tagging.TagSet, b.ko.Spec.Tagging.TagSet)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Versioning, b.ko.Spec.Versioning) {
		delta.Add("Spec.Versioning", a.ko.Spec.Versioning, b.ko.Spec.Versioning)
	} else if a.ko.Spec.Versioning != nil && b.ko.Spec.Versioning != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.Versioning.Status, b.ko.Spec.Versioning.Status) {
			delta.Add("Spec.Versioning.Status", a.ko.Spec.Versioning.Status, b.ko.Spec.Versioning.Status)
		} else if a.ko.Spec.Versioning.Status != nil && b.ko.Spec.Versioning.Status != nil {
			if *a.ko.Spec.Versioning.Status != *b.ko.Spec.Versioning.Status {
				delta.Add("Spec.Versioning.Status", a.ko.Spec.Versioning.Status, b.ko.Spec.Versioning.Status)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Website, b.ko.Spec.Website) {
		delta.Add("Spec.Website", a.ko.Spec.Website, b.ko.Spec.Website)
	} else if a.ko.Spec.Website != nil && b.ko.Spec.Website != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.Website.ErrorDocument, b.ko.Spec.Website.ErrorDocument) {
			delta.Add("Spec.Website.ErrorDocument", a.ko.Spec.Website.ErrorDocument, b.ko.Spec.Website.ErrorDocument)
		} else if a.ko.Spec.Website.ErrorDocument != nil && b.ko.Spec.Website.ErrorDocument != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Website.ErrorDocument.Key, b.ko.Spec.Website.ErrorDocument.Key) {
				delta.Add("Spec.Website.ErrorDocument.Key", a.ko.Spec.Website.ErrorDocument.Key, b.ko.Spec.Website.ErrorDocument.Key)
			} else if a.ko.Spec.Website.ErrorDocument.Key != nil && b.ko.Spec.Website.ErrorDocument.Key != nil {
				if *a.ko.Spec.Website.ErrorDocument.Key != *b.ko.Spec.Website.ErrorDocument.Key {
					delta.Add("Spec.Website.ErrorDocument.Key", a.ko.Spec.Website.ErrorDocument.Key, b.ko.Spec.Website.ErrorDocument.Key)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Website.IndexDocument, b.ko.Spec.Website.IndexDocument) {
			delta.Add("Spec.Website.IndexDocument", a.ko.Spec.Website.IndexDocument, b.ko.Spec.Website.IndexDocument)
		} else if a.ko.Spec.Website.IndexDocument != nil && b.ko.Spec.Website.IndexDocument != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Website.IndexDocument.Suffix, b.ko.Spec.Website.IndexDocument.Suffix) {
				delta.Add("Spec.Website.IndexDocument.Suffix", a.ko.Spec.Website.IndexDocument.Suffix, b.ko.Spec.Website.IndexDocument.Suffix)
			} else if a.ko.Spec.Website.IndexDocument.Suffix != nil && b.ko.Spec.Website.IndexDocument.Suffix != nil {
				if *a.ko.Spec.Website.IndexDocument.Suffix != *b.ko.Spec.Website.IndexDocument.Suffix {
					delta.Add("Spec.Website.IndexDocument.Suffix", a.ko.Spec.Website.IndexDocument.Suffix, b.ko.Spec.Website.IndexDocument.Suffix)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Website.RedirectAllRequestsTo, b.ko.Spec.Website.RedirectAllRequestsTo) {
			delta.Add("Spec.Website.RedirectAllRequestsTo", a.ko.Spec.Website.RedirectAllRequestsTo, b.ko.Spec.Website.RedirectAllRequestsTo)
		} else if a.ko.Spec.Website.RedirectAllRequestsTo != nil && b.ko.Spec.Website.RedirectAllRequestsTo != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.Website.RedirectAllRequestsTo.HostName, b.ko.Spec.Website.RedirectAllRequestsTo.HostName) {
				delta.Add("Spec.Website.RedirectAllRequestsTo.HostName", a.ko.Spec.Website.RedirectAllRequestsTo.HostName, b.ko.Spec.Website.RedirectAllRequestsTo.HostName)
			} else if a.ko.Spec.Website.RedirectAllRequestsTo.HostName != nil && b.ko.Spec.Website.RedirectAllRequestsTo.HostName != nil {
				if *a.ko.Spec.Website.RedirectAllRequestsTo.HostName != *b.ko.Spec.Website.RedirectAllRequestsTo.HostName {
					delta.Add("Spec.Website.RedirectAllRequestsTo.HostName", a.ko.Spec.Website.RedirectAllRequestsTo.HostName, b.ko.Spec.Website.RedirectAllRequestsTo.HostName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.Website.RedirectAllRequestsTo.Protocol, b.ko.Spec.Website.RedirectAllRequestsTo.Protocol) {
				delta.Add("Spec.Website.RedirectAllRequestsTo.Protocol", a.ko.Spec.Website.RedirectAllRequestsTo.Protocol, b.ko.Spec.Website.RedirectAllRequestsTo.Protocol)
			} else if a.ko.Spec.Website.RedirectAllRequestsTo.Protocol != nil && b.ko.Spec.Website.RedirectAllRequestsTo.Protocol != nil {
				if *a.ko.Spec.Website.RedirectAllRequestsTo.Protocol != *b.ko.Spec.Website.RedirectAllRequestsTo.Protocol {
					delta.Add("Spec.Website.RedirectAllRequestsTo.Protocol", a.ko.Spec.Website.RedirectAllRequestsTo.Protocol, b.ko.Spec.Website.RedirectAllRequestsTo.Protocol)
				}
			}
		}
		if !reflect.DeepEqual(a.ko.Spec.Website.RoutingRules, b.ko.Spec.Website.RoutingRules) {
			delta.Add("Spec.Website.RoutingRules", a.ko.Spec.Website.RoutingRules, b.ko.Spec.Website.RoutingRules)
		}
	}

	return delta
}
